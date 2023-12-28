package amapiv1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
)

// AnimeMovieServer serves gRPC requests for Anime Movie endpoints.
type AnimeMovieServer struct {
	ampbv1.UnimplementedAnimeMovieServiceServer
	config     utils.Config
	gojo       db.Gojo
	tokenMaker token.Maker
	ping       *ping.PingSystem
}

func NewAnimeMovieServer(config utils.Config, gojo db.Gojo, tokenMaker token.Maker, ping *ping.PingSystem) *AnimeMovieServer {
	return &AnimeMovieServer{
		config:     config,
		gojo:       gojo,
		tokenMaker: tokenMaker,
		ping:       ping,
	}
}
