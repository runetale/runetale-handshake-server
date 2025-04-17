package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/runetale/client-go/runetale/runetale/v1/negotiation"
	"github.com/runetale/client-go/runetale/runetale/v1/ping"
	"github.com/runetale/client-go/runetale/runetale/v1/rtc"
	"github.com/runetale/runetale-handshake-server/domain/entity"
	server "github.com/runetale/runetale-handshake-server/grpc_server"
	"github.com/runetale/runetale-handshake-server/utility"
	"google.golang.org/grpc"
	health "google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

var (
	env    *entity.Env
	err    error
	logger *utility.Logger
)

func init() {
	// get env
	env = entity.NewEnv()

	// initialize logger
	logger, err = utility.NewLogger(env.GetLogFile(), env.LogFmt, env.LogLevel)
	if err != nil {
		log.Fatalf("failed to initialize logger: %v", err)
		panic(err)
	}
}

func main() {
	pCache := entity.NewPeerCache()

	s := server.NewServer(pCache, env, logger)

	opts := server.NewGrpcServerOption()
	grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer())
	grpcServer := grpc.NewServer(opts...)

	negotiation.RegisterNegotiationServiceServer(grpcServer, s.NegotiatinServer)
	rtc.RegisterRtcServiceServer(grpcServer, s.RtcServer)
	ping.RegisterPingServiceServer(grpcServer, s.PingServer)
	healthpb.RegisterHealthServer(grpcServer, health.NewServer())

	logger.Info(fmt.Sprintf("starting server with :%s", env.Port))
	logger.Info(fmt.Sprintf("stun server => %s", env.StunConfig.URL))
	logger.Info(fmt.Sprintf("turn server => %s", env.TurnConfig.URL))

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", env.Port))
	if err != nil {
		logger.Error("failed to listen: %v", err)
	}

	reflection.Register(grpcServer)
	go func() {
		if err = grpcServer.Serve(lis); err != nil {
			logger.Error("failed to serve grpc server: %v.", err)
		}
	}()

	stop := make(chan struct{})
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c,
			syscall.SIGTERM,
			syscall.SIGINT,
		)
		select {
		case <-c:
			close(stop)
		}
	}()
	<-stop
	logger.Info("terminated server")
	grpcServer.Stop()
}
