package asapiv1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	aspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/aspb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
)

// AnimeSerieServer serves gRPC requests for Anime Serie endpoints.
type AnimeSerieServer struct {
	aspbv1.UnimplementedAnimeSerieServiceServer
	config     utils.Config
	gojo       db.Gojo
	tokenMaker token.Maker
	ping       *ping.PingSystem
}

func NewAnimeSerieServer(config utils.Config, gojo db.Gojo, tokenMaker token.Maker, ping *ping.PingSystem) *AnimeSerieServer {
	return &AnimeSerieServer{
		config:     config,
		gojo:       gojo,
		tokenMaker: tokenMaker,
		ping:       ping,
	}
}
