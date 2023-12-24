package animeSerie

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/go-redis/cache/v9"
)

// AnimeSerieServer serves gRPC requests for Anime Serie endpoints.
type AnimeSerieServer struct {
	config     utils.Config
	gojo       db.Gojo
	tokenMaker token.Maker
	cache      *cache.Cache
}

func NewAnimeSerieServer(config utils.Config, gojo db.Gojo, tokenMaker token.Maker, cache *cache.Cache) *AnimeSerieServer {
	return &AnimeSerieServer{
		config:     config,
		gojo:       gojo,
		tokenMaker: tokenMaker,
		cache:      cache,
	}
}
