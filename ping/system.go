package ping

import (
	"fmt"
	"time"

	"github.com/dj-yacine-flutter/gojo/conf"
	"github.com/go-redis/cache/v9"
	"github.com/redis/go-redis/v9"
)

// PingSystem is a cache system using redis.
type PingSystem struct {
	config *conf.DataEnv
	cache  *cache.Cache
}

func NewPingSystem(host string, config *conf.DataEnv) *PingSystem {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server01": host,
		},
	})

	cache := cache.New(&cache.Options{
		Redis:      ring,
		LocalCache: cache.NewTinyLFU(90, time.Minute),
	})

	fmt.Printf("\u001b[38;5;50mSTART CACHE SERVER -AT- %s\u001b[0m\n", host)

	return &PingSystem{
		config: config,
		cache:  cache,
	}
}
