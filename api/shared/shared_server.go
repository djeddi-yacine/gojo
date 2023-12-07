package shared

import (
	"github.com/dj-yacine-flutter/gojo/pb"
	"github.com/dj-yacine-flutter/gojo/token"
)

// SharedServer manage our gRPC requests.
type SharedServer struct {
	pb.UnimplementedGojoServer
	tokenMaker token.Maker
}

func NewSharedServer(tokenMaker token.Maker) *SharedServer {
	return &SharedServer{
		tokenMaker: tokenMaker,
	}
}
