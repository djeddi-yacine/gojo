package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	v1 "github.com/dj-yacine-flutter/gojo/api/v1"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/dj-yacine-flutter/gojo/worker"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
	"golang.org/x/sync/errgroup"
)

func main() {
	var err error

	config, err := utils.LoadConfig(".", "gojo")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config file")
	}

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM, syscall.SIGINT)
	defer stop()

	conn, err := pgxpool.New(ctx, config.DBSource)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot connect to the DB")
	}

	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		log.Fatal().Err(err).Msg("cannot create token maker")
	}

	fmt.Printf("\u001b[38;5;125m\u001b[48;5;0m%s\u001b[0m\n", fmt.Sprintln(`
                                            
     ██████╗  ██████╗      ██╗ ██████╗      
    ██╔════╝ ██╔═══██╗     ██║██╔═══██╗     
    ██║  ███╗██║   ██║     ██║██║   ██║     
    ██║   ██║██║   ██║██   ██║██║   ██║     
    ╚██████╔╝╚██████╔╝╚█████╔╝╚██████╔╝     
     ╚═════╝  ╚═════╝  ╚════╝  ╚═════╝      
                                            `))

	client := utils.MeiliSearch(config)

	ping := ping.NewPingSystem(config)

	gojo := db.NewGojo(conn, ping)

	redisOpt := asynq.RedisClientOpt{
		Addr: config.RedisQueueAddress,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpt)

	waitGroup, ctx := errgroup.WithContext(ctx)

	worker.Start(ctx, waitGroup, config, redisOpt, gojo)

	v1.Start(ctx, waitGroup, config, gojo, tokenMaker, taskDistributor, ping, client)

	err = waitGroup.Wait()
	if err != nil {
		log.Fatal().Err(err).Msg("error from wait group")
	}
}
