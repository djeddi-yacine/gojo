package v1

import (
	"context"
	"fmt"
	"net"
	"net/http"

	"github.com/dj-yacine-flutter/gojo/api"
	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	amapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime/movie"
	asapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime/serie"
	nfapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/info"
	usapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/user"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	_ "github.com/dj-yacine-flutter/gojo/doc/v1/statik"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	apbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	uspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/uspb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/dj-yacine-flutter/gojo/worker"
	runtime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/meilisearch/meilisearch-go"
	sk "github.com/rakyll/statik/fs"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func Start(config utils.Config, gojo db.Gojo, tokenMaker token.Maker, taskDistributor worker.TaskDistributor, ping *ping.PingSystem, client *meilisearch.Client) {
	amx, err := utils.CreatedIndex(client, utils.AnimeMovieV1)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	asx, err := utils.CreatedIndex(client, utils.AnimeSeasonV1)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	ussvc := usapiv1.NewUserServer(config, gojo, tokenMaker, taskDistributor)
	nfsvc := nfapiv1.NewInfoServer(gojo, tokenMaker)
	asvc := av1.NewAnimeServer(gojo, tokenMaker)
	amsvc := amapiv1.NewAnimeMovieServer(gojo, tokenMaker, ping, amx)
	assvc := asapiv1.NewAnimeSerieServer(gojo, tokenMaker, ping, asx)

	go startGatewayApi(config, ussvc, nfsvc, asvc, amsvc, assvc)
	startGRPCApi(config, ussvc, nfsvc, asvc, amsvc, assvc)
}

func startGRPCApi(
	config utils.Config,
	ussvc *usapiv1.UserServer,
	nfsvc *nfapiv1.InfoServer,
	asvc *av1.AnimeServer,
	amsvc *amapiv1.AnimeMovieServer,
	assvc *asapiv1.AnimeSerieServer,
) {
	var err error

	gprcLogger := grpc.UnaryInterceptor(api.GrpcLogger)
	server := grpc.NewServer(gprcLogger)

	uspbv1.RegisterUserServiceServer(server, ussvc)
	nfpbv1.RegisterInfoServiceServer(server, nfsvc)
	apbv1.RegisterAnimeServiceServer(server, asvc)
	ampbv1.RegisterAnimeMovieServiceServer(server, amsvc)
	aspbv1.RegisterAnimeSerieServiceServer(server, assvc)

	reflection.Register(server)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create gRPC listener")
	}

	fmt.Printf("\u001b[38;5;50m\u001b[48;5;0m- START gRPC server -AT- %s\u001b[0m\n", listener.Addr().String())

	err = server.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start gRPC server")
	}
}

func startGatewayApi(
	config utils.Config,
	ussvc *usapiv1.UserServer,
	nfsvc *nfapiv1.InfoServer,
	asvc *av1.AnimeServer,
	amsvc *amapiv1.AnimeMovieServer,
	assvc *asapiv1.AnimeSerieServer,
) {
	var err error

	httpMux := http.NewServeMux()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	grpcMux := runtime.NewServeMux()

	err = uspbv1.RegisterUserServiceHandlerServer(ctx, grpcMux, ussvc)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register Gateway server for User Service v1")
	}

	err = nfpbv1.RegisterInfoServiceHandlerServer(ctx, grpcMux, nfsvc)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register Gateway server for Info Service v1")
	}

	err = apbv1.RegisterAnimeServiceHandlerServer(ctx, grpcMux, asvc)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register Gateway server for Anime Service v1")
	}

	err = ampbv1.RegisterAnimeMovieServiceHandlerServer(ctx, grpcMux, amsvc)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register Gateway server for Anime Movie Service v1")
	}

	err = aspbv1.RegisterAnimeSerieServiceHandlerServer(ctx, grpcMux, assvc)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register Gateway server for Anime Serie Service v1")
	}

	httpMux.Handle("/v1/", http.StripPrefix("/v1", grpcMux))

	statikFS, err := sk.New()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create statik files for swagger v1")
	}

	httpMux.Handle("/v1/swagger/", http.StripPrefix("/v1/swagger/", http.FileServer(statikFS)))

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create Gateway listener")
	}

	fmt.Printf("\u001b[38;5;50m\u001b[48;5;0m- START HTTP server -AT- %s\u001b[0m\n", listener.Addr().String())

	err = http.Serve(listener, httpMux)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start the Gateway server")
	}
}
