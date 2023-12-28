package ping

import (
	"fmt"
	"time"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

// PingSystem is a cache system using redis.
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

	cache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(90, time.Minute),
	})

	fmt.Printf("\u001b[38;5;50m\u001b[48;5;0m- START REDIS CACHE SERVER -AT- %s\u001b[0m\n", config.RedisCacheAddress)

	return &PingSystem{
		config: config,
		cache:  cache,
	}
}
