package user

import (
	db "github.com/dj-yacine-flutter/gojo/db/database"
	"github.com/dj-yacine-flutter/gojo/token"
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/dj-yacine-flutter/gojo/worker"
)

// UserServer serves gRPC requests for User endpoints.
type UserServer struct {
	config          utils.Config
	gojo            db.Gojo
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

func NewUserServer(config utils.Config, gojo db.Gojo, tokenMaker token.Maker, taskDistributor worker.TaskDistributor) *UserServer {
	return &UserServer{
		config:          config,
		gojo:            gojo,
		tokenMaker:      tokenMaker,
		taskDistributor: taskDistributor,
	}
}
