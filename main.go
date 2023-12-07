package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	rt "runtime"

	"github.com/dj-yacine-flutter/gojo/api"
	"github.com/dj-yacine-flutter/gojo/api/shared"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	_ "github.com/dj-yacine-flutter/gojo/doc/statik"
	"github.com/dj-yacine-flutter/gojo/mail"
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/dj-yacine-flutter/gojo/worker"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	sk "github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func init() {
	// Set GOMAXPROCS to the number of available CPU cores
	rt.GOMAXPROCS(rt.NumCPU())
}

func main() {
	config, err := utils.LoadConfig(".", "gojo")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config file")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to the DB")
	}

	gojo := db.NewGojo(conn)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisAddress,
	}

	taskDistributot := worker.NewRedisTaskDistributor(redisOpt)

	go runTaskProcessor(config, redisOpt, gojo)
	go runGatewayServer(config, gojo, taskDistributot)

	fmt.Printf(`
	 ██████╗  ██████╗      ██╗ ██████╗ 
	██╔════╝ ██╔═══██╗     ██║██╔═══██╗
	██║  ███╗██║   ██║     ██║██║   ██║
	██║   ██║██║   ██║██   ██║██║   ██║
	╚██████╔╝╚██████╔╝╚█████╔╝╚██████╔╝
	 ╚═════╝  ╚═════╝  ╚════╝  ╚═════╝ 
									   
`)
	runGRPCServer(config, gojo, taskDistributot)
}

func runTaskProcessor(config utils.Config, redisOpt asynq.RedisClientOpt, gojo db.Gojo) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, gojo, mailer)

	log.Info().Msg("start task processor")
	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}

func runGRPCServer(config utils.Config, gojo db.Gojo, taskDistrinbutor worker.TaskDistributor) {
	server, err := api.NewServer(config, gojo, taskDistrinbutor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create GRPC server")
	}

	gprcLogger := grpc.UnaryInterceptor(shared.GrpcLogger)
	grpcServer := grpc.NewServer(gprcLogger)
	pb.RegisterGojoServer(grpcServer, server)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", config.GRPCServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create gRPC listener")
	}

	log.Info().Msgf("start gRPC server at %s", listener.Addr().String())
	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start the gRPC server")
	}
}

func runGatewayServer(config utils.Config, gojo db.Gojo, taskDistrinbutor worker.TaskDistributor) {
	server, err := api.NewServer(config, gojo, taskDistrinbutor)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create Gateway server")
	}

	grpcMux := runtime.NewServeMux()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = pb.RegisterGojoHandlerServer(ctx, grpcMux, server)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot register Gateway server")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	statikFS, err := sk.New()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create statik")
	}
	swaggerHandler := http.StripPrefix("/swagger/", http.FileServer(statikFS))
	mux.Handle("/swagger/", swaggerHandler)

	listener, err := net.Listen("tcp", config.HTTPServerAddress)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create Gateway listener")
	}

	log.Info().Msgf("start HTTP gateway server at %s", listener.Addr().String())

	handler := shared.HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start the Gateway server")
	}
}
