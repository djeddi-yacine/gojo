package amapiv1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	ampbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/ampb"
	"github.com/dj-yacine-flutter/gojo/ping"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/meilisearch/meilisearch-go"
)

// AnimeMovieServer serves gRPC requests for Anime Movie endpoints.
type AnimeMovieServer struct {
	ampbv1.UnimplementedAnimeMovieServiceServer
	gojo        db.Gojo
	tokenMaker  token.Maker
	ping        *ping.PingSystem
	meilisearch *meilisearch.Index
}

func NewAnimeMovieServer(gojo db.Gojo, tokenMaker token.Maker, ping *ping.PingSystem, index *meilisearch.Index) *AnimeMovieServer {
	return &AnimeMovieServer{
		gojo:        gojo,
		tokenMaker:  tokenMaker,
		ping:        ping,
		meilisearch: index,
	}
}
