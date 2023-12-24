package api

import (
	"fmt"

	animeMovie "github.com/dj-yacine-flutter/gojo/api/anime/movie"
	animeSerie "github.com/dj-yacine-flutter/gojo/api/anime/serie"
	"github.com/dj-yacine-flutter/gojo/api/info"
	"github.com/dj-yacine-flutter/gojo/api/user"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/dj-yacine-flutter/gojo/worker"
	"github.com/go-redis/cache/v9"
)

// Server serves gRPC requests for our Gojo service.
type Server struct {
	*user.UserServer
	*info.InfoServer
	*animeSerie.AnimeSerieServer
	*animeMovie.AnimeMovieServer
}

// NewServer creates a new gRPC server.
func NewServer(config utils.Config, gojo db.Gojo, taskDistributor worker.TaskDistributor, cache *cache.Cache) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		UserServer:       user.NewUserServer(config, gojo, tokenMaker, taskDistributor),
		InfoServer:       info.NewInfoServer(gojo, tokenMaker),
		AnimeSerieServer: animeSerie.NewAnimeSerieServer(config, gojo, tokenMaker, cache),
		AnimeMovieServer: animeMovie.NewAnimeMovieServer(config, gojo, tokenMaker, cache),
	}

	return server, nil
}
