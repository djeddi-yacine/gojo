package conf

import "time"

type DataEnv struct {
	TokenSymmetricKey    string        // TOKEN_SYMMETRIC_KEY
	MeiliSearchMasterKey string        // MEILISEATCH_MASTER_KEY
	AccessTokenDuration  time.Duration // ACCESS_TOKEN_DURATION
	RefreshTokenDuration time.Duration // REFRESH_TOKEN_DURATION
	CacheRepetition      uint8         // CACHE_REPETITION
	CacheKeyDuration     time.Duration // CACHE_KEY_DURATION
	CacheCountDuration   time.Duration // CACHE_COUNT_DURATION
}
