package v1

import (
	"context"
	"fmt"
	"net/http"

	usapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/user"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	uspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/uspb"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/dj-yacine-flutter/gojo/worker"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func StartGRPCApi(server *grpc.Server, config utils.Config, gojo db.Gojo, taskDistributor worker.TaskDistributor) error {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return fmt.Errorf("cannot create token maker: %w", err)
	}

	user := usapiv1.NewUserServer(config, gojo, tokenMaker, taskDistributor)

	uspbv1.RegisterUserServiceServer(server, user)

	return nil
}

func StartGatewayApi(httpMux *http.ServeMux, config utils.Config, gojo db.Gojo, taskDistributor worker.TaskDistributor) error {
	var err error

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return fmt.Errorf("cannot create token maker: %w", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcMux := runtime.NewServeMux()

	user := usapiv1.NewUserServer(config, gojo, tokenMaker, taskDistributor)

	err = uspbv1.RegisterUserServiceHandlerServer(ctx, grpcMux, user)
	if err != nil {
		return fmt.Errorf("cannot register Gateway server for User Service v1: %w", err)
	}

	vMux := http.NewServeMux()
	vMux.Handle("/v1/", http.StripPrefix("/v1", grpcMux))

	// Use the custom ServeMux
	httpMux.Handle("/", vMux)

	return nil
}
