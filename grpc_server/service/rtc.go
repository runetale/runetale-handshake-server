package service

import (
	"context"

	"github.com/runetale/client-go/runetale/runetale/v1/rtc"
	"github.com/runetale/runetale-handshake-server/domain/entity"
	"google.golang.org/protobuf/types/known/emptypb"
)

type RtcServiceImpl interface {
	GetStunTurnConfig(ctx context.Context, req *emptypb.Empty) (*rtc.GetStunTurnConfigResponse, error)
}

type RtcService struct {
	env *entity.Env
}

func NewRtcService(e *entity.Env) RtcServiceImpl {
	return &RtcService{
		env: e,
	}
}

func (s *RtcService) AuthFuncOverride(ctx context.Context, fullMethodName string) (context.Context, error) {
	return ctx, nil
}

// TODO: add a few more stun turns
func (s *RtcService) GetStunTurnConfig(ctx context.Context, req *emptypb.Empty) (*rtc.GetStunTurnConfigResponse, error) {
	th := &rtc.TurnHost{
		Url:      s.env.TurnConfig.URL,
		Username: s.env.TurnConfig.Username,
		Password: s.env.TurnConfig.Password,
	}

	sh := &rtc.StunHost{
		Url:      s.env.StunConfig.URL,
		Username: s.env.StunConfig.Username,
		Password: s.env.StunConfig.Password,
	}

	rtcconf := &rtc.RtcConfig{
		TurnHost: th,
		StunHost: sh,
	}

	return &rtc.GetStunTurnConfigResponse{
		RtcConfig: rtcconf,
	}, nil
}
