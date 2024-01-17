package utils

import (
	"fmt"

	"github.com/meilisearch/meilisearch-go"
	"github.com/rs/zerolog/log"
)

const (
	AnimeMovieV1  = "AnimeMovieV1"
	AnimeSeasonV1 = "AnimeSeasonV1"
)

type Document struct {
	ID     int64    `json:"ID"`
	Titles []string `json:"titles"`
}

func MeiliSearch(config Config) *meilisearch.Client {
	meiliClient := meilisearch.NewClient(meilisearch.ClientConfig{
		Host: "http://" + config.MeilisearchAddress,
	})

	fmt.Printf("\u001b[38;5;50m\u001b[48;5;0m- START MEILISEATCH -AT- %s\u001b[0m\n", config.MeilisearchAddress)

	return meiliClient
}

func CreatedIndex(client *meilisearch.Client, ID string) (*meilisearch.Index, error) {
	var err error

	it, err := client.CreateIndex(&meilisearch.IndexConfig{
		Uid:        ID,
		PrimaryKey: "ID",
	})
	if err != nil {
		log.Fatal().Err(err).Msgf("cannot create index in ID: %s", ID)
	}

	_, err = client.WaitForTask(it.TaskUID)
	if err != nil {
		log.Fatal().Err(err).Msgf("cannot wait for task in ID: %s", ID)
	}

	at, err := client.Index(ID).UpdateFilterableAttributes(&[]string{"ID"})
	if err != nil {
		log.Fatal().Err(err).Msgf("cannot create [Attributes] index in ID: %s", ID)
	}

	_, err = client.WaitForTask(at.TaskUID)
	if err != nil {
		log.Fatal().Err(err).Msgf("cannot wait for [Attributes] task in ID: %s", ID)
	}

	return client.Index(ID), nil
}
