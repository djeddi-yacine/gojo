package conf

import "net/url"

type ServerEnv struct {
	HTTPAddress        *url.URL // HTTP_ADDRESS
	GRPCAddress        *url.URL // GRPC_ADDRESS
	RedisQueueAddress  *url.URL // REDIS_QUEUE_ADDRESS
	RedisCacheAddress  *url.URL // REDIS_CACHE_ADDRESS
	MeilisearchAddress *url.URL // MEILISEATCH_ADDRESS
}
