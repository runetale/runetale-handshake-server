package service

import (
	"context"

	"github.com/runetale/client-go/runetale/runetale/v1/negotiation"
	"github.com/runetale/runetale-handshake-server/domain/entity"
	"github.com/runetale/runetale-handshake-server/infura/di"
	"github.com/runetale/runetale-handshake-server/utility"
	"google.golang.org/protobuf/types/known/emptypb"
)

type NegotiationServiceImpl interface {
	Offer(ctx context.Context, req *negotiation.HandshakeRequest) (*emptypb.Empty, error)
	Answer(ctx context.Context, req *negotiation.HandshakeRequest) (*emptypb.Empty, error)
	Candidate(ctx context.Context, req *negotiation.CandidateRequest) (*emptypb.Empty, error)
	Connect(stream negotiation.NegotiationService_ConnectServer) error
}

type NegotiationService struct {
	pCache *entity.PeerCache
	logger *utility.Logger
}

func NewNegotiationService(p *entity.PeerCache, logger *utility.Logger) NegotiationServiceImpl {
	return &NegotiationService{
		pCache: p,
		logger: logger,
	}
}

func (s *NegotiationService) Offer(ctx context.Context, req *negotiation.HandshakeRequest) (*emptypb.Empty, error) {
	negInteractor := di.InitialNegotiationInteractor(s.pCache, s.logger)
	return negInteractor.Offer(req)
}

func (s *NegotiationService) Answer(ctx context.Context, req *negotiation.HandshakeRequest) (*emptypb.Empty, error) {
	negInteractor := di.InitialNegotiationInteractor(s.pCache, s.logger)
	return negInteractor.Answer(req)
}

func (s *NegotiationService) Candidate(ctx context.Context, req *negotiation.CandidateRequest) (*emptypb.Empty, error) {
	negInteractor := di.InitialNegotiationInteractor(s.pCache, s.logger)
	return negInteractor.Candidate(req)
}

func (s *NegotiationService) Connect(stream negotiation.NegotiationService_ConnectServer) error {
	negoInteractor := di.InitialNegotiationInteractor(s.pCache, s.logger)

	_, err := negoInteractor.Connect(stream)
	if err != nil {
		return err
	}

	return nil
}
