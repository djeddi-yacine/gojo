package nfapiv1

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	nfpbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/nfpb"
	"github.com/dj-yacine-flutter/gojo/token"
)

// InfoServer serves gRPC requests for Info endpoints.
type InfoServer struct {
	nfpbv1.UnimplementedInfoServiceServer
	gojo       db.Gojo
	tokenMaker token.Maker
}

func NewInfoServer(gojo db.Gojo, tokenMaker token.Maker) *InfoServer {
	return &InfoServer{
		gojo:       gojo,
		tokenMaker: tokenMaker,
	}
}
