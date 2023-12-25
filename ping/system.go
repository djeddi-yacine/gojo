package ping

import (
	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/go-redis/cache/v9"
)

// PingSystem serves gRPC requests for Anime Movie endpoints.
type PingSystem struct {
	config utils.Config
	cache  *cache.Cache
}

func NewPingSystem(config utils.Config, cache *cache.Cache) *PingSystem {
	return &PingSystem{
		config: config,
		cache:  cache,
	}
}
