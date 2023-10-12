package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"

	db "github.com/dj-yacine-flutter/gojo/db/database"
	_ "github.com/dj-yacine-flutter/gojo/doc/statik"
	"github.com/dj-yacine-flutter/gojo/gapi"
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/pgxpool"
	sk "github.com/rakyll/statik/fs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	config, err := utils.LoadConfig(".")
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

	runDBMigration(config.MigrationURL, config.DBSource)

	gojo := db.NewGojo(conn)

	go runGatewayServer(config, gojo)
	fmt.Printf(`
	++=========================================================++
	||  ++=================================================++  ||
	||  ||:::::::::::::::::::::::::::::::::::::::::::::::::||  ||
	||  ||::::'######::::'#######::::::::'##::'#######:::::||  ||
	||  ||:::'##... ##::'##.... ##::::::: ##:'##.... ##::::||  ||
	||  ||::: ##:::..::: ##:::: ##::::::: ##: ##:::: ##::::||  ||
	||  ||::: ##::'####: ##:::: ##::::::: ##: ##:::: ##::::||  ||
	||  ||::: ##::: ##:: ##:::: ##:'##::: ##: ##:::: ##::::||  ||
	||  ||::: ##::: ##:: ##:::: ##: ##::: ##: ##:::: ##::::||  ||
	||  ||:::. ######:::. #######::. ######::. #######:::::||  ||
	||  ||::::......:::::.......::::......::::.......::::::||  ||
	||  ||:::::::::::::::::::::::::::::::::::::::::::::::::||  ||
	||  ++=================================================++  ||
	++=========================================================++

`)
	runGRPCServer(config, gojo)
}

func runDBMigration(migrationURL string, dbSource string) {
	migration, err := migrate.New(migrationURL, dbSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create new migrate instance")
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal().Err(err).Msg("failed to run migrate up")
	}

	log.Info().Msg("db migrated successfully")
}

func runGRPCServer(config utils.Config, gojo db.Gojo) {
	server, err := gapi.NewServer(config, gojo)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create GRPC server")
	}

	gprcLogger := grpc.UnaryInterceptor(gapi.GrpcLogger)
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

func runGatewayServer(config utils.Config, gojo db.Gojo) {
	server, err := gapi.NewServer(config, gojo)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create Gateway server")
	}

	/* 	jsonOption := runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
		MarshalOptions: protojson.MarshalOptions{
			UseProtoNames: true,
		},
		UnmarshalOptions: protojson.UnmarshalOptions{
			DiscardUnknown: true,
		},
	}) */
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

	handler := gapi.HttpLogger(mux)
	err = http.Serve(listener, handler)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start the Gateway server")
	}
}
