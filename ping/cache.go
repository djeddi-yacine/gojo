package ping

import (
	"context"
	"fmt"

	"github.com/go-redis/cache/v9"
	"github.com/rs/zerolog/log"
)

const (
	ANIME_MOVIE   = "AM"
	ANIME_SERIE   = "AS"
	ANIME_SEASON  = "AX"
	ANIME_EPISODE = "AE"
)

type KeyGenrator interface {
	Key() string
	Count() string
}

type CacheKey struct {
	ID     int64
	Target string
	key    string
}

func (x *CacheKey) Key() string {
	return x.key
}

func (x *CacheKey) Count() string {
	return x.key + ":COUNT"
}

func (system *PingSystem) Handle(ctx context.Context, gen KeyGenrator, value interface{}, fn func() error) error {
	var err error
	if err = system.cache.GetSkippingLocalCache(ctx, gen.Key(), value); err != nil {
		if err = fn(); err != nil {
			return err
		}

		var target uint8
		_ = system.cache.GetSkippingLocalCache(ctx, gen.Count(), &target)
		if target < system.config.CacheRepetition {
			target++
			if err = system.cache.Set(&cache.Item{
				Ctx:            ctx,
				Key:            gen.Count(),
				Value:          &target,
				TTL:            system.config.CacheCountDuration,
				SkipLocalCache: true,
			}); err != nil {
				log.Err(err)
				return nil
			}
		}

		if target >= system.config.CacheRepetition {
			if err = system.cache.Set(&cache.Item{
				Ctx:            ctx,
				Key:            gen.Key(),
				Value:          value,
				TTL:            system.config.CacheKeyDuration,
				SkipLocalCache: true,
			}); err != nil {
				log.Err(err)
				return nil
			}
			_ = system.cache.Delete(ctx, gen.Count())
		}

		log.Debug().
			Str("key", gen.Key()).
			Bool("found", false).
			Uint8("count", target).
			Msg("cache system")
	} else {
		log.Debug().
			Str("key", gen.Key()).
			Bool("found", true).
			Msg("cache system")
	}

	return err
}

func (x *CacheKey) Main() KeyGenrator {
	return &CacheKey{
		key: fmt.Sprintf("%s:%d", x.Target, x.ID),
	}
}

func (x *CacheKey) Meta(language uint32) KeyGenrator {
	return &CacheKey{
		key: fmt.Sprintf("%s:MTD:%d:LNG:%d", x.Target, language, x.ID),
	}
}

func (x *CacheKey) Studio() KeyGenrator {
	return &CacheKey{
		key: fmt.Sprintf("%s:STD:%d", x.Target, x.ID),
	}
}

func (x *CacheKey) Genre() KeyGenrator {
	return &CacheKey{
		key: fmt.Sprintf("%s:GNR:%d", x.Target, x.ID),
	}
}

func (x *CacheKey) Resources() KeyGenrator {
	return &CacheKey{
		key: fmt.Sprintf("%s:RSC:%d", x.Target, x.ID),
	}
}

func (x *CacheKey) Links() KeyGenrator {
	return &CacheKey{
		key: fmt.Sprintf("%s:LNK:%d", x.Target, x.ID),
	}
}

func (x *CacheKey) Server() KeyGenrator {
	return &CacheKey{
		key: fmt.Sprintf("%s:SRV:%d", x.Target, x.ID),
	}
}

func (x *CacheKey) Sub() KeyGenrator {
	return &CacheKey{
		key: fmt.Sprintf("%s:SUB:%d", x.Target, x.ID),
	}
}

func (x *CacheKey) Dub() KeyGenrator {
	return &CacheKey{
		key: fmt.Sprintf("%s:DUB:%d", x.Target, x.ID),
	}
}

func (x *CacheKey) Tags() KeyGenrator {
	return &CacheKey{
		key: fmt.Sprintf("%s:TAG:%d", x.Target, x.ID),
	}
}

func (x *CacheKey) Images() KeyGenrator {
	return &CacheKey{
		key: fmt.Sprintf("%s:IMG:%d", x.Target, x.ID),
	}
}

func (x *CacheKey) Trailers() KeyGenrator {
	return &CacheKey{
		key: fmt.Sprintf("%s:TRL:%d", x.Target, x.ID),
	}
}
