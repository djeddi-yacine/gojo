package ping

import (
	"context"
	"reflect"

	"github.com/go-redis/cache/v9"
	"github.com/rs/zerolog/log"
)

const (
	AnimeMovie   = "AM"
	AnimeSerie   = "AS"
	AnimeSeason  = "AX"
	AnimeEpisode = "AE"
)

type KeyGenrator interface {
	Key() string
	Count() string
}

type PingKey string

func (x PingKey) Key() string {
	return string(x)
}

func (x PingKey) Count() string {
	return string(x + "-COUNT")
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
		system.cache.GetSkippingLocalCache(ctx, gen.Count(), &target)
		switch {
		case target < system.config.CacheRepetition:
			target++
			if err = system.cache.Set(&cache.Item{
				Ctx:            ctx,
				Key:            gen.Count(),
				Value:          &target,
				TTL:            system.config.CacheCountDuration,
				SkipLocalCache: true,
			}); err != nil {
				log.Err(err).Msg("Error setting cache for target")
				return nil
			}
		case target >= system.config.CacheRepetition:
			if err = system.cache.Set(&cache.Item{
				Ctx:            ctx,
				Key:            gen.Key(),
				Value:          value,
				TTL:            system.config.CacheKeyDuration,
				SkipLocalCache: true,
			}); err != nil {
				log.Err(err).Msg("Error setting cache for key")
				return nil
			}
			if system.cache.Exists(ctx, gen.Count()) {
				if err := system.cache.Delete(ctx, gen.Count()); err != nil {
					log.Err(err).Msg("Error deleting cache for count")
					return nil
				}
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

	return nil
}
