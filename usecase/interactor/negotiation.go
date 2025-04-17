package interactor

import (
	"fmt"
	"io"

	"github.com/runetale/client-go/runetale/runetale/v1/negotiation"
	"github.com/runetale/runetale-handshake-server/domain/entity"
	repository "github.com/runetale/runetale-handshake-server/domain/interface"
	infura "github.com/runetale/runetale-handshake-server/infura/interface"
	"github.com/runetale/runetale-handshake-server/utility"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NegotiationInteractorImpl struct {
	peerRepository repository.PeerRepository
	logger         *utility.Logger
}

func NewNegotiationInteractorImpl(
	pcr repository.PeerRepository,
	logger *utility.Logger,
) infura.NegotiationInteractor {
	return &NegotiationInteractorImpl{
		peerRepository: pcr,
		logger:         logger,
	}
}

// Called when the runetale agent connects to the Stream of negotiation.Negotiation_ConnectClient
func (i *NegotiationInteractorImpl) Connect(streamServer negotiation.NegotiationService_ConnectServer) (*entity.PeerStream, error) {
	header := map[string][]string{"Content-Type": {"application/grpc"}}
	streamServer.SetHeader(header)
	// receive node-key from peer connected to StreamConnectClient
	nk, err := getNodeKeyFromCtx(streamServer.Context())
	if err != nil {
		return nil, err
	}

	_, err = getWgPubKeyFromCtx(streamServer.Context())
	if err != nil {
		return nil, err
	}

	host, err := getHostFromCtx(streamServer.Context())
	if err != nil {
		return nil, err
	}

	os, err := getOSFromCtx(streamServer.Context())
	if err != nil {
		return nil, err
	}

	distro, err := getDistroFromCtx(streamServer.Context())
	if err != nil {
		return nil, err
	}

	i.logger.Debug(fmt.Sprintf("connected node. [Host: %s, OS: %s, Distro: %s]", os, host, distro))

	// cache the stream and nodekey of the peer in local in-memory
	peerStream := entity.NewPeerStream(nk, streamServer)
	err = i.peerRepository.Create(peerStream)
	if err != nil {
		return peerStream, err
	}

	defer func() {
		i.logger.Debug(fmt.Sprintf("disconnect from [%s]. OS: %s, HOST: %s", nk, os, host))
		i.Disconnect(peerStream)
	}()

	// continue to receive messages from stream client(remote peer)
	// a message is sent from another remote peer
	for {
		// when the Offer, Answer, or Candidate RPC is called,
		msg, err := streamServer.Recv()
		if err == io.EOF {
			i.logger.Error(fmt.Sprintf("return to EOF, received by [%s]", nk), err)
			break
		}
		if err != nil {
			i.logger.Error(fmt.Sprintf("close to signal connect %s", err.Error()), err)
			return peerStream, err
		}

		// find another peer by remote peer node key
		if dstPeer, found := i.peerRepository.FindByClientNodeKey(msg.GetDstNodeKey()); found {
			// sent to the part of the StartConnectClient stream where the destination peer is doing the `Recv`
			err := dstPeer.Stream.Send(msg)
			if err != nil {
				i.logger.Error(fmt.Sprintf("error: can not send dst peer: %s. from: %s", msg.GetDstNodeKey(), nk), err)
				return peerStream, err
			}
			i.logger.Debug(fmt.Sprintf("send to %s stream from: %s", nk, msg.GetDstNodeKey()))
		} else {
			i.logger.Error(fmt.Sprintf("can not found the dstPeer => [%s]", msg.GetDstNodeKey()), err)
		}
	}
	<-streamServer.Context().Done()

	return peerStream, streamServer.Context().Err()
}

func (i *NegotiationInteractorImpl) Disconnect(peer *entity.PeerStream) error {
	return i.peerRepository.Delete(peer)
}

func (i *NegotiationInteractorImpl) Offer(req *negotiation.HandshakeRequest) (*emptypb.Empty, error) {
	var err error

	// confirm that the peer to which the message is to be sent is registered
	if !i.peerRepository.IsCreated(req.GetSrcNodeKey()) {
		return nil, entity.ErrPeerNotRegister
	}

	if dstPeer, found := i.peerRepository.FindByClientNodeKey(req.GetDstNodeKey()); found {
		err = dstPeer.Stream.Send(
			&negotiation.NegotiationRequest{
				Type:        negotiation.NegotiationType_OFFER,
				DstNodeKey:  req.GetSrcNodeKey(),
				DstWgPubKey: req.GetWgPubKey(),
				UFlag:       req.GetUFlag(),
				Pwd:         req.GetPwd(),
			},
		)
		if err != nil {
			return nil, err
		}

		i.logger.Debug(fmt.Sprintf("[OFFER] found the dstPeer => [%s]", dstPeer.ClientNodeKey))
	} else {
		i.logger.Error(fmt.Sprintf("[OFFER] can not found the dstPeer => [%s], it appears that the other party has not launched the runetale client", req.GetDstNodeKey()), err)
	}

	return &emptypb.Empty{}, nil
}

func (i *NegotiationInteractorImpl) Answer(req *negotiation.HandshakeRequest) (*emptypb.Empty, error) {
	var err error
	// confirm that the peer to which the message is to be sent is registered
	if !i.peerRepository.IsCreated(req.GetSrcNodeKey()) {
		return nil, entity.ErrPeerNotRegister
	}

	if dstPeer, found := i.peerRepository.FindByClientNodeKey(req.GetDstNodeKey()); found {
		err = dstPeer.Stream.Send(
			&negotiation.NegotiationRequest{
				Type:        negotiation.NegotiationType_ANSWER,
				DstNodeKey:  req.GetSrcNodeKey(),
				DstWgPubKey: req.GetWgPubKey(),
				UFlag:       req.GetUFlag(),
				Pwd:         req.GetPwd(),
			},
		)
		if err != nil {
			return nil, err
		}

		i.logger.Debug(fmt.Sprintf("[ANSWER] found the [%s] from [%s]", dstPeer.ClientNodeKey, req.GetSrcNodeKey()))
		i.logger.Debug(fmt.Sprintf("uflag: %s, pwd: %s", req.GetUFlag(), req.GetPwd()))
	} else {
		i.logger.Error(fmt.Sprintf("[Answer] can not found the dstPeer => [%s]", dstPeer.ClientNodeKey), err)
	}

	return &emptypb.Empty{}, nil
}

func (i *NegotiationInteractorImpl) Candidate(req *negotiation.CandidateRequest) (*emptypb.Empty, error) {
	var err error
	// confirm that the peer to which the message is to be sent is registered
	if !i.peerRepository.IsCreated(req.GetSrcNodeKey()) {
		return nil, entity.ErrPeerNotRegister
	}

	if dstPeer, found := i.peerRepository.FindByClientNodeKey(req.GetDstNodeKey()); found {
		err = dstPeer.Stream.Send(
			&negotiation.NegotiationRequest{
				Type:        negotiation.NegotiationType_CANDIDATE,
				DstNodeKey:  req.GetSrcNodeKey(),
				DstWgPubKey: req.GetWgPubKey(),
				Candidate:   req.GetCandidate(),
			},
		)
		if err != nil {
			return nil, err
		}
		i.logger.Debug(fmt.Sprintf("[CANDIDATE] found the [%s] from [%s]", dstPeer.ClientNodeKey, req.GetSrcNodeKey()))
	} else {
		i.logger.Error(fmt.Sprintf("[CANDIDATE] can not found the dstPeer => [%s]", dstPeer.ClientNodeKey), err)
	}

	return &emptypb.Empty{}, nil
}
