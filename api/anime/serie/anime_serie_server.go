package animeSerie

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
)

// AnimeSerieServer serves gRPC requests for Anime Serie endpoints.
type AnimeSerieServer struct {
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
