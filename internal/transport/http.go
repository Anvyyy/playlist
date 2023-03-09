package transport

import (
	"context"
	"log"
	"net/http"

	"github.com/Anvyyy/playlist/pkg"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func RunHTTPserver(ctx context.Context, gRPCport, httpPort string) error{
	conn, err := grpc.Dial(
		gRPCport,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	gateway := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = pkg.RegisterMusicServiceHandlerFromEndpoint(ctx, gateway, gRPCport, opts)
	if err != nil {
		return err
	}

	return http.ListenAndServe(httpPort, gateway)
}