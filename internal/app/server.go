package app

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/Anvyyy/playlist/internal/transport"
	"github.com/Anvyyy/playlist/pkg"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
)

func Run(gRPCport, httpPort string) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	listener, err := net.Listen("tcp", gRPCport)
	if err != nil {
		return fmt.Errorf("error lister grpc port")
	}

	server := grpc.NewServer(
		grpc.UnaryInterceptor(unaryInterceptor),
	)
	playlist := transport.NewGrpcTransport()
	pkg.RegisterMusicServiceServer(server, playlist)
	go func() {
		if err = server.Serve(listener); err != nil {
			log.Err(err).Msgf("error serve")
		}
	}()

	transport.RunHTTPserver(ctx, gRPCport, httpPort)

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	defer signal.Stop(ch)

	<-ch
	server.GracefulStop()

	return nil
}

func unaryInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
	log.Info().Msgf("unary interceptor: %s", info.FullMethod)

	return handler(ctx, req)
}