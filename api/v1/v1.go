package v1

import (
	"context"
	"fmt"
	"net/http"

	nfapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/info"
	usapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/user"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
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

	info := nfapiv1.NewInfoServer(gojo, tokenMaker)
	nfpbv1.RegisterInfoServiceServer(server, info)

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

	info := nfapiv1.NewInfoServer(gojo, tokenMaker)
	err = nfpbv1.RegisterInfoServiceHandlerServer(ctx, grpcMux, info)
	if err != nil {
		return fmt.Errorf("cannot register Gateway server for Info Service v1: %w", err)
	}

	vMux := http.NewServeMux()
	vMux.Handle("/v1/", http.StripPrefix("/v1", grpcMux))

	// Use the custom ServeMux
	httpMux.Handle("/", vMux)

	return nil
}
