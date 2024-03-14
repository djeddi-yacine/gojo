package usapiv1

import (
	"github.com/dj-yacine-flutter/gojo/conf"
	db "github.com/dj-yacine-flutter/gojo/db/database"
	uspbv1 "github.com/dj-yacine-flutter/gojo/pb/v1/uspb"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/worker"
)

// UserServer serves gRPC requests for User endpoints.
type UserServer struct {
	uspbv1.UnimplementedUserServiceServer
	config          *conf.DataEnv
	gojo            db.Gojo
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewUserServer(config *conf.DataEnv, gojo db.Gojo, tokenMaker token.Maker, taskDistributor worker.TaskDistributor) *UserServer {
	return &UserServer{
		config:          config,
		gojo:            gojo,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}
}
