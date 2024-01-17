package asapiv1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/meilisearch/meilisearch-go"
)

// AnimeSerieServer serves gRPC requests for Anime Serie endpoints.
type AnimeSerieServer struct {
	aspbv1.UnimplementedAnimeSerieServiceServer
	config      utils.Config
	gojo        db.Gojo
	tokenMaker  token.Maker
	ping        *ping.PingSystem
	meilisearch *meilisearch.Index
}

func NewAnimeSerieServer(config utils.Config, gojo db.Gojo, tokenMaker token.Maker, ping *ping.PingSystem, index *meilisearch.Index) *AnimeSerieServer {
	return &AnimeSerieServer{
		config:      config,
		gojo:        gojo,
		tokenMaker:  tokenMaker,
		ping:        ping,
		meilisearch: index,
	}
}
