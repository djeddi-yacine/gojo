package v1

import (
	"context"
	"fmt"
	"net/http"

	amapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime/movie"
	asapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime/serie"
	nfapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/info"
	usapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/user"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	uspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/uspb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/dj-yacine-flutter/gojo/worker"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func StartGRPCApi(server *grpc.Server, config utils.Config, gojo db.Gojo, tokenMaker token.Maker, taskDistributor worker.TaskDistributor, ping *ping.PingSystem) error {
	ussvc := usapiv1.NewUserServer(config, gojo, tokenMaker, taskDistributor)
	uspbv1.RegisterUserServiceServer(server, ussvc)

	nfsvc := nfapiv1.NewInfoServer(gojo, tokenMaker)
	nfpbv1.RegisterInfoServiceServer(server, nfsvc)

	amsvc := amapiv1.NewAnimeMovieServer(config, gojo, tokenMaker, ping)
	ampbv1.RegisterAnimeMovieServiceServer(server, amsvc)

	assvc := asapiv1.NewAnimeSerieServer(config, gojo, tokenMaker, ping)
	aspbv1.RegisterAnimeSerieServiceServer(server, assvc)

	return nil
}

func StartGatewayApi(httpMux *http.ServeMux, config utils.Config, gojo db.Gojo, tokenMaker token.Maker, taskDistributor worker.TaskDistributor, ping *ping.PingSystem) error {
	var err error

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcMux := runtime.NewServeMux()

	ussvc := usapiv1.NewUserServer(config, gojo, tokenMaker, taskDistributor)

	err = uspbv1.RegisterUserServiceHandlerServer(ctx, grpcMux, ussvc)
	if err != nil {
		return fmt.Errorf("cannot register Gateway server for User Service v1: %w", err)
	}

	nfsvc := nfapiv1.NewInfoServer(gojo, tokenMaker)
	err = nfpbv1.RegisterInfoServiceHandlerServer(ctx, grpcMux, nfsvc)
	if err != nil {
		return fmt.Errorf("cannot register Gateway server for Info Service v1: %w", err)
	}

	amsvc := amapiv1.NewAnimeMovieServer(config, gojo, tokenMaker, ping)
	err = ampbv1.RegisterAnimeMovieServiceHandlerServer(ctx, grpcMux, amsvc)
	if err != nil {
		return fmt.Errorf("cannot register Gateway server for Anime Movie Service v1: %w", err)
	}

	assvc := asapiv1.NewAnimeSerieServer(config, gojo, tokenMaker, ping)
	err = aspbv1.RegisterAnimeSerieServiceHandlerServer(ctx, grpcMux, assvc)
	if err != nil {
		return fmt.Errorf("cannot register Gateway server for Anime Serie Service v1: %w", err)
	}

	// Use the custom ServeMux
	httpMux.Handle("/v1/", http.StripPrefix("/v1", grpcMux))

	return nil
}
