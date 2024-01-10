package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	"github.com/dj-yacine-flutter/gojo/api"
	v1 "github.com/dj-yacine-flutter/gojo/api/v1"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/mail"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/dj-yacine-flutter/gojo/worker"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	var err error

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

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create token maker")
	}

	gojo := db.NewGojo(conn)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisQueueAddress,
	}

	fmt.Printf("\u001b[38;5;125m\u001b[48;5;0m%s\u001b[0m\n", fmt.Sprintln(`
                                            
     ██████╗  ██████╗      ██╗ ██████╗      
    ██╔════╝ ██╔═══██╗     ██║██╔═══██╗     
    ██║  ███╗██║   ██║     ██║██║   ██║     
    ██║   ██║██║   ██║██   ██║██║   ██║     
    ╚██████╔╝╚██████╔╝╚█████╔╝╚██████╔╝     
     ╚═════╝  ╚═════╝  ╚════╝  ╚═════╝      
                                            `))

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)
	ping := ping.NewPingSystem(config)

	go queue(config, redisOpt, gojo)
	go v1Http(config, gojo, tokenMaker, taskDistributor, ping)
	v1Grpc(config, gojo, tokenMaker, taskDistributor, ping)
}

func v1Http(config utils.Config, gojo db.Gojo, tokenMaker token.Maker, taskDistributor worker.TaskDistributor, ping *ping.PingSystem) {
	var err error

	httpMux := http.NewServeMux()
	err = v1.StartGatewayApi(httpMux, config, gojo, tokenMaker, taskDistributor, ping)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

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

func v1Grpc(config utils.Config, gojo db.Gojo, tokenMaker token.Maker, taskDistributor worker.TaskDistributor, ping *ping.PingSystem) {
	var err error

	gprcLogger := grpc.UnaryInterceptor(api.GrpcLogger)
	server := grpc.NewServer(gprcLogger)

	err = v1.StartGRPCApi(server, config, gojo, tokenMaker, taskDistributor, ping)
	if err != nil {
		log.Fatal().Err(err).Msg(err.Error())
	}

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

func queue(config utils.Config, redisOpt asynq.RedisClientOpt, gojo db.Gojo) {
	mailer := mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, gojo, mailer)

	fmt.Printf("\u001b[38;5;50m\u001b[48;5;0m- START REDIS TASK PROCESSOR -AT- %s\u001b[0m\n", config.RedisQueueAddress)

	err := taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to start task processor")
	}
}
