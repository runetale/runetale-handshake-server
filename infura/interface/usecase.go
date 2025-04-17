package infura

import (
	"github.com/runetale/client-go/runetale/runetale/v1/negotiation"
	"github.com/runetale/runetale-handshake-server/domain/entity"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NegotiationInteractor interface {
	Offer(req *negotiation.HandshakeRequest) (*emptypb.Empty, error)
	Answer(req *negotiation.HandshakeRequest) (*emptypb.Empty, error)
	Candidate(req *negotiation.CandidateRequest) (*emptypb.Empty, error)

	Connect(streamServer negotiation.NegotiationService_ConnectServer) (*entity.PeerStream, error)
	Disconnect(peer *entity.PeerStream) error
}
