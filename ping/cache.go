package ping

import (
	"context"
	"fmt"
	"reflect"

	"github.com/go-redis/cache/v9"
	"github.com/rs/zerolog/log"
)

type CacheKey struct {
	ID      int64
	Target  string
	Version rune
}

const (
	AnimeMovie        = "AM"
	AnimeSerie        = "AS"
	AnimeSeason       = "AX"
	AnimeEpisode      = "AE"
	V1           rune = '1'
)

type KeyGenrator interface {
	Key() string
	Count() string
}

type PingKey struct {
	key string
}

func (x *PingKey) Key() string {
	return x.key
}

func (x *PingKey) Count() string {
	return x.key + ":COUNT"
}

func (system *PingSystem) Handle(ctx context.Context, gen KeyGenrator, value interface{}, fn func() error) error {
	var err error
	if err = system.cache.GetSkippingLocalCache(ctx, gen.Key(), value); err != nil {
		if err = fn(); err != nil {
			return err
		}

		v := reflect.ValueOf(value)
		for v.Kind() == reflect.Ptr {
			if v.IsNil() || v.Elem().IsZero() {
				log.Warn().
					Str("key", gen.Key()).
					Msg("value is nil, not storing in cache")
				return nil
			}
			v = v.Elem()
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
			if system.cache.Exists(ctx, gen.Count()) {
				_ = system.cache.Delete(ctx, gen.Count())
			}
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
	return &PingKey{
		key: fmt.Sprintf("V%c:%s:ID:%d", x.Version, x.Target, x.ID),
	}
}

func (x *CacheKey) Meta(language uint32) KeyGenrator {
	return &PingKey{
		key: fmt.Sprintf("V%c:%s:MTD:%d:LNG:%d", x.Version, x.Target, language, x.ID),
	}
}

func (x *CacheKey) Studio() KeyGenrator {
	return &PingKey{
		key: fmt.Sprintf("V%c:%s:STD:%d", x.Version, x.Target, x.ID),
	}
}

func (x *CacheKey) Genre() KeyGenrator {
	return &PingKey{
		key: fmt.Sprintf("V%c:%s:GNR:%d", x.Version, x.Target, x.ID),
	}
}

func (x *CacheKey) Resources() KeyGenrator {
	return &PingKey{
		key: fmt.Sprintf("V%c:%s:RSC:%d", x.Version, x.Target, x.ID),
	}
}

func (x *CacheKey) Links() KeyGenrator {
	return &PingKey{
		key: fmt.Sprintf("V%c:%s:LNK:%d", x.Version, x.Target, x.ID),
	}
}

func (x *CacheKey) Server() KeyGenrator {
	return &PingKey{
		key: fmt.Sprintf("V%c:%s:SRV:%d", x.Version, x.Target, x.ID),
	}
}

func (x *CacheKey) Sub() KeyGenrator {
	return &PingKey{
		key: fmt.Sprintf("V%c:%s:SUB:%d", x.Version, x.Target, x.ID),
	}
}

func (x *CacheKey) Dub() KeyGenrator {
	return &PingKey{
		key: fmt.Sprintf("V%c:%s:DUB:%d", x.Version, x.Target, x.ID),
	}
}

func (x *CacheKey) Tags() KeyGenrator {
	return &PingKey{
		key: fmt.Sprintf("V%c:%s:TAG:%d", x.Version, x.Target, x.ID),
	}
}

func (x *CacheKey) Images() KeyGenrator {
	return &PingKey{
		key: fmt.Sprintf("V%c:%s:IMG:%d", x.Version, x.Target, x.ID),
	}
}

func (x *CacheKey) Trailers() KeyGenrator {
	return &PingKey{
		key: fmt.Sprintf("V%c:%s:TRL:%d", x.Version, x.Target, x.ID),
	}
}
