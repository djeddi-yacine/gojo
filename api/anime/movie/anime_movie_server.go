package animeMovie

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/go-redis/cache/v9"
)

// AnimeMovieServer serves gRPC requests for Anime Movie endpoints.
type AnimeMovieServer struct {
	config     utils.Config
	gojo       db.Gojo
	tokenMaker token.Maker
	cache      *cache.Cache
}

func NewAnimeMovieServer(config utils.Config, gojo db.Gojo, tokenMaker token.Maker, cache *cache.Cache) *AnimeMovieServer {
	return &AnimeMovieServer{
		config:     config,
		gojo:       gojo,
		tokenMaker: tokenMaker,
		cache:      cache,
	}
}
