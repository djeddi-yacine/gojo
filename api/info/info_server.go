package info

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/token"
)

// InfoServer serves gRPC requests for Info endpoints.
type InfoServer struct {
	pb.UnimplementedGojoServer
	gojo       db.Gojo
	tokenMaker token.Maker
}

func NewInfoServer(gojo db.Gojo, tokenMaker token.Maker) *InfoServer {
	return &InfoServer{
		gojo:       gojo,
		tokenMaker: tokenMaker,
	}
}
