package service

import (
	"context"

	"github.com/runetale/client-go/runetale/runetale/v1/ping"
	"google.golang.org/protobuf/types/known/emptypb"
)

type PingServiceImpl interface {
	Ping(ctx context.Context, req *emptypb.Empty) (*ping.PingResponse, error)
}

type PingService struct {
}

func NewPingService() PingServiceImpl {
	return &PingService{}
}

func (s *PingService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

func (s *PingService) Ping(ctx context.Context, req *emptypb.Empty) (*ping.PingResponse, error) {
	return &ping.PingResponse{}, nil
}
