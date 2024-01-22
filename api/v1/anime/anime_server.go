package av1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	apbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/apb"
	"github.com/dj-yacine-flutter/gojo/token"
)

// AnimeServer serves gRPC requests for Anime  endpoints.
type AnimeServer struct {
	apbv1.UnimplementedAnimeServiceServer
	gojo       db.Gojo
	tokenMaker token.Maker
}

func NewAnimeServer(gojo db.Gojo, tokenMaker token.Maker) *AnimeServer {
	return &AnimeServer{
		gojo:       gojo,
		tokenMaker: tokenMaker,
	}
}
