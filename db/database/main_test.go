package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/dj-yacine-flutter/gojo/conf"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testGojo Gojo

func TestMain(m *testing.M) {
	config, err := conf.Load("../..", "example.gojo")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DB.URL())
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	ping := ping.NewPingSystem(config.Server.RedisCacheAddress.Host, config.Data)

	testGojo = NewGojo(connPool, ping)
	os.Exit(m.Run())
}
