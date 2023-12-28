package ping

import (
	"fmt"
	"time"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

// PingSystem serves gRPC requests for Anime Movie endpoints.
type PingSystem struct {
	config utils.Config
	cache  *cache.Cache
}

func NewPingSystem(config utils.Config) *PingSystem {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server01": config.RedisCacheAddress,
		},
	})

	fmt.Printf("\u001b[38;5;50m\u001b[48;5;0m- START REDIS CACHE SERVER -AT- %s\u001b[0m\n", config.RedisCacheAddress)

	cache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(90, time.Minute),
	})

	return &PingSystem{
		config: config,
		cache:  cache,
	}
}
