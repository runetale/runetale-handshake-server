package grpcserver

import (
	"github.com/runetale/runetale-handshake-server/domain/entity"
	"github.com/runetale/runetale-handshake-server/grpc_server/service"
	"github.com/runetale/runetale-handshake-server/utility"
)

type Server struct {
	NegotiatinServer service.NegotiationServiceImpl
	RtcServer        service.RtcServiceImpl
	PingServer       service.PingServiceImpl
}

func NewServer(p *entity.PeerCache, env *entity.Env, logger *utility.Logger) *Server {
	return &Server{
		NegotiatinServer: service.NewNegotiationService(p, logger),
		RtcServer:        service.NewRtcService(env),
		PingServer:       service.NewPingService(),
	}
}
