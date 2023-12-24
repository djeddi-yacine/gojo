package animeMovie

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/rs/zerolog/log"
)

type KeyGenerator interface {
	Code() string
	Key() string
}

type AnimeKey struct {
	id  int64
	key string
}

func (x *AnimeKey) Code() string {
	return fmt.Sprintf("anime_movie:::%d", x.id)
}

func (x *AnimeKey) Key() string {
	return x.key
}

func (server *AnimeMovieServer) do(ctx context.Context, gen KeyGenerator, value interface{}, fn func() error) error {
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

func (x *AnimeKey) Anime() KeyGenerator {
	return &AnimeKey{
		id:  x.id,
		key: fmt.Sprintf("anime_movie:anime:%d", x.id),
	}
}

func (x *AnimeKey) Meta(language uint32) KeyGenerator {
	return &AnimeKey{
		id:  x.id,
		key: fmt.Sprintf("anime_movie:meta:%d:%d", language, x.id),
	}
}

func (x *AnimeKey) Studio() KeyGenerator {
	return &AnimeKey{
		id:  x.id,
		key: fmt.Sprintf("anime_movie:studio:%d", x.id),
	}
}

func (x *AnimeKey) Genre() KeyGenerator {
	return &AnimeKey{
		id:  x.id,
		key: fmt.Sprintf("anime_movie:genre:%d", x.id),
	}
}

func (x *AnimeKey) Resources() KeyGenerator {
	return &AnimeKey{
		id:  x.id,
		key: fmt.Sprintf("anime_movie:resources:%d", x.id),
	}
}

func (x *AnimeKey) Links() KeyGenerator {
	return &AnimeKey{
		id:  x.id,
		key: fmt.Sprintf("anime_movie:links:%d", x.id),
	}
}

func (x *AnimeKey) Server() KeyGenerator {
	return &AnimeKey{
		id:  x.id,
		key: fmt.Sprintf("anime_movie:server:%d", x.id),
	}
}

func (x *AnimeKey) Sub() KeyGenerator {
	return &AnimeKey{
		id:  x.id,
		key: fmt.Sprintf("anime_movie:sub:%d", x.id),
	}
}

func (x *AnimeKey) Dub() KeyGenerator {
	return &AnimeKey{
		id:  x.id,
		key: fmt.Sprintf("anime_movie:dub:%d", x.id),
	}
}

func (x *AnimeKey) Tags() KeyGenerator {
	return &AnimeKey{
		id:  x.id,
		key: fmt.Sprintf("anime_movie:tags:%d", x.id),
	}
}

func (x *AnimeKey) Images() KeyGenerator {
	return &AnimeKey{
		id:  x.id,
		key: fmt.Sprintf("anime_movie:images:%d", x.id),
	}
}

func (x *AnimeKey) Trailers() KeyGenerator {
	return &AnimeKey{
		id:  x.id,
		key: fmt.Sprintf("anime_movie:trailers:%d", x.id),
	}
}
