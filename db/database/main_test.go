package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgx/v5/pgxpool"
)

var testGojo Gojo

func TestMain(m *testing.M) {
	config, err := utils.LoadConfig("../..", "example.gojo")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}

	connPool, err := pgxpool.New(context.Background(), config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testGojo = NewGojo(connPool)
	os.Exit(m.Run())
}
