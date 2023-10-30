package api

import (
	"fmt"

	animeMovie "github.com/dj-yacine-flutter/gojo/api/anime/movie"
	animeSerie "github.com/dj-yacine-flutter/gojo/api/anime/serie"
	"github.com/dj-yacine-flutter/gojo/api/info"
	"github.com/dj-yacine-flutter/gojo/api/shared"
	"github.com/dj-yacine-flutter/gojo/api/user"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/dj-yacine-flutter/gojo/worker"
)

// Server serves gRPC requests for our Gojo service.
type Server struct {
	*shared.SharedServer
	*user.UserServer
	*info.InfoServer
	*animeSerie.AnimeSerieServer
	*animeMovie.AnimeMovieServer
}

// NewServer creates a new gRPC server.
func NewServer(config utils.Config, gojo db.Gojo, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		SharedServer:     shared.NewSharedServer(tokenMaker),
		UserServer:       user.NewUserServer(config, gojo, tokenMaker, taskDistributor),
		InfoServer:       info.NewInfoServer(gojo, tokenMaker),
		AnimeSerieServer: animeSerie.NewAnimeSerieServer(config, gojo, tokenMaker, taskDistributor),
		AnimeMovieServer: animeMovie.NewAnimeMovieServer(config, gojo, tokenMaker, taskDistributor),
	}

	return server, nil
}
