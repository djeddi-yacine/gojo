package v1

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"

	"github.com/dj-yacine-flutter/gojo/api"
	av1 "github.com/dj-yacine-flutter/gojo/api/v1/anime"
	amapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime/movie"
	asapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/anime/serie"
	nfapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/info"
	usapiv1 "github.com/dj-yacine-flutter/gojo/api/v1/user"
	"github.com/dj-yacine-flutter/gojo/conf"
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
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
)

func Start(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config *conf.Config,
	gojo db.Gojo,
	tokenMaker token.Maker,
	taskDistributor worker.TaskDistributor,
	ping *ping.PingSystem,
	client *meilisearch.Client,
) {
	amx, err := utils.CreatedIndex(client, utils.AnimeMovieV1)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	asx, err := utils.CreatedIndex(client, utils.AnimeSeasonV1)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

	ussvc := usapiv1.NewUserServer(config.Data, gojo, tokenMaker, taskDistributor)
	nfsvc := nfapiv1.NewInfoServer(gojo, tokenMaker)
	asvc := av1.NewAnimeServer(gojo, tokenMaker)
	amsvc := amapiv1.NewAnimeMovieServer(gojo, tokenMaker, ping, amx)
	assvc := asapiv1.NewAnimeSerieServer(gojo, tokenMaker, ping, asx)

	startGatewayApi(ctx, waitGroup, config.Server, ussvc, nfsvc, asvc, amsvc, assvc)
	startGRPCApi(ctx, waitGroup, config.Server, ussvc, nfsvc, asvc, amsvc, assvc)
}

func startGRPCApi(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config *conf.ServerEnv,
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

	listener, err := net.Listen("tcp", config.GRPCAddress.Host)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create gRPC listener")
	}

	waitGroup.Go(func() error {
		fmt.Printf("\u001b[38;5;50mSTART gRPC SERVER -AT- %s\u001b[0m\n", config.GRPCAddress.Host)

		err = server.Serve(listener)
		if err != nil {
			if errors.Is(err, grpc.ErrServerStopped) {
				return nil
			}
			log.Error().Err(err).Msg("gRPC server failed to serve")
			return err
		}

		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		fmt.Printf("\u001b[38;5;266mSHUTDOWN gRPC SERVER ...\u001b[0m\n")

		server.GracefulStop()
		fmt.Printf("\u001b[38;5;196mgRPC SERVER IS STOPPED\u001b[0m\n")

		return nil
	})

}

func startGatewayApi(
	ctx context.Context,
	waitGroup *errgroup.Group,
	config *conf.ServerEnv,
	ussvc *usapiv1.UserServer,
	nfsvc *nfapiv1.InfoServer,
	asvc *av1.AnimeServer,
	amsvc *amapiv1.AnimeMovieServer,
	assvc *asapiv1.AnimeSerieServer,
) {
	var err error

	grpcMux := runtime.NewServeMux(runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	}))

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

	httpMux := http.NewServeMux()
	httpMux.Handle("/v1/", http.StripPrefix("/v1", grpcMux))

	statikFS, err := sk.New()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create statik files for swagger v1")
	}

	httpMux.Handle("/v1/swagger/", http.StripPrefix("/v1/swagger/", http.FileServer(statikFS)))

	httpServer := &http.Server{
		Handler: httpMux,
		Addr:    config.HTTPAddress.Host,
	}

	waitGroup.Go(func() error {
		fmt.Printf("\u001b[38;5;50mSTART HTTP SERVER -AT- %s\u001b[0m\n", config.HTTPAddress.Host)

		err = httpServer.ListenAndServe()
		if err != nil {
			if errors.Is(err, http.ErrServerClosed) {
				return nil
			}
			log.Error().Err(err).Msg("HTTP server failed to serve")
			return err
		}
		return nil
	})

	waitGroup.Go(func() error {
		<-ctx.Done()
		fmt.Printf("\u001b[38;5;266mSHUTDOWN HTTP SERVER ...\u001b[0m\n")

		err := httpServer.Shutdown(context.Background())
		if err != nil {
			log.Error().Err(err).Msg("failed to shutdown HTTP server")
			return err
		}

		fmt.Printf("\u001b[38;5;196mHTTP SERVER IS STOPPED\u001b[0m\n")
		return nil
	})
}
