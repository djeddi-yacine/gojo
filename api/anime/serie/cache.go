package animeSerie

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/rs/zerolog/log"
)

const (
	SERIE_KEY   = "AS"
	SEASON_KEY  = "AX"
	EPISODE_KEY = "AE"
)

type KeyGenerator interface {
	Code() string
	Key() string
}

type CacheKey struct {
	id     int64
	target string
	key    string
}

func (x *CacheKey) Code() string {
	return fmt.Sprintf("%s:::%d", x.target, x.id)
}

func (x *CacheKey) Key() string {
	return x.key
}

func (server *AnimeSerieServer) do(ctx context.Context, gen KeyGenerator, value interface{}, fn func() error) error {
	var err error
	if err = server.cache.Get(ctx, gen.Key(), value); err != nil {
		if err = fn(); err != nil {
			return err
		}

		var rip uint8
		if err = server.cache.GetSkippingLocalCache(ctx, gen.Code(), &rip); err != nil {
			if err = server.cache.Set(&cache.Item{
				Ctx:   ctx,
				Key:   gen.Code(),
				Value: &rip,
				TTL:   6 * time.Hour,
				SetNX: true,
			}); err != nil {
				log.Err(err)
				return nil
			}
		}

		if rip >= server.config.CacheRepetition {
			if err = server.cache.Set(&cache.Item{
				Ctx:   ctx,
				Key:   gen.Key(),
				Value: value,
				TTL:   12 * time.Hour,
				SetNX: true,
			}); err != nil {
				log.Err(err)
				return nil
			}
		} else {
			rip++
			if err = server.cache.Set(&cache.Item{
				Ctx:            ctx,
				Key:            gen.Code(),
				Value:          &rip,
				TTL:            6 * time.Hour,
				SetXX:          true,
				SkipLocalCache: true,
			}); err != nil {
				log.Err(err)
				return nil
			}
		}

		log.Info().Msgf("See %s in cache %d time", gen.Key(), rip)
	} else {
		log.Info().Msgf("Get %s from cache successfully", gen.Key())
	}

	return err
}

func (x *CacheKey) Anime() KeyGenerator {
	return &CacheKey{
		id:     x.id,
		target: x.target,
		key:    fmt.Sprintf("%s:anime:%d", x.target, x.id),
	}
}

func (x *CacheKey) Meta(language uint32) KeyGenerator {
	return &CacheKey{
		id:     x.id,
		target: x.target,
		key:    fmt.Sprintf("%s:meta:%d:%d", x.target, language, x.id),
	}
}

func (x *CacheKey) Studio() KeyGenerator {
	return &CacheKey{
		id:     x.id,
		target: x.target,
		key:    fmt.Sprintf("%s:studio:%d", x.target, x.id),
	}
}

func (x *CacheKey) Genre() KeyGenerator {
	return &CacheKey{
		id:     x.id,
		target: x.target,
		key:    fmt.Sprintf("%s:genre:%d", x.target, x.id),
	}
}

func (x *CacheKey) Resources() KeyGenerator {
	return &CacheKey{
		id:     x.id,
		target: x.target,
		key:    fmt.Sprintf("%s:resources:%d", x.target, x.id),
	}
}

func (x *CacheKey) Links() KeyGenerator {
	return &CacheKey{
		id:     x.id,
		target: x.target,
		key:    fmt.Sprintf("%s:links:%d", x.target, x.id),
	}
}

func (x *CacheKey) Server() KeyGenerator {
	return &CacheKey{
		id:     x.id,
		target: x.target,
		key:    fmt.Sprintf("%s:server:%d", x.target, x.id),
	}
}

func (x *CacheKey) Sub() KeyGenerator {
	return &CacheKey{
		id:     x.id,
		target: x.target,
		key:    fmt.Sprintf("%s:sub:%d", x.target, x.id),
	}
}

func (x *CacheKey) Dub() KeyGenerator {
	return &CacheKey{
		id:     x.id,
		target: x.target,
		key:    fmt.Sprintf("%s:dub:%d", x.target, x.id),
	}
}

func (x *CacheKey) Tags() KeyGenerator {
	return &CacheKey{
		id:     x.id,
		target: x.target,
		key:    fmt.Sprintf("%s:tags:%d", x.target, x.id),
	}
}

func (x *CacheKey) Images() KeyGenerator {
	return &CacheKey{
		id:     x.id,
		target: x.target,
		key:    fmt.Sprintf("%s:images:%d", x.target, x.id),
	}
}

func (x *CacheKey) Trailers() KeyGenerator {
	return &CacheKey{
		id:     x.id,
		target: x.target,
		key:    fmt.Sprintf("%s:trailers:%d", x.target, x.id),
	}
}
