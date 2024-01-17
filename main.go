package main

import (
	"context"
	"fmt"

	v1 "github.com/dj-yacine-flutter/gojo/api/v1"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/mail"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/dj-yacine-flutter/gojo/worker"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog/log"
)

func main() {
	var err error

	config, err := utils.LoadConfig(".", "gojo")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config file")
	}

	conn, err := pgxpool.New(context.Background(), config.DBSource)
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
	taskProcessor := worker.NewRedisTaskProcessor(redisOpt, gojo, mail.NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword))

	err = taskProcessor.Start()
	if err != nil {
		log.Fatal().Err(err).Msg("cannot start task processor")
	}

	v1.Start(config, gojo, tokenMaker, taskDistributor, ping, client)
}
