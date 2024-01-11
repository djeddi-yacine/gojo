package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/ping"
)

func (gojo *SQLGojo) ListAnimeTagsTx(ctx context.Context, arg []int64) ([]AnimeTag, error) {
	var err error
	var result []AnimeTag

	err = gojo.execTx(ctx, func(q *Queries) error {
		var cache ping.SegmentKey
		result = make([]AnimeTag, len(arg))

		for i, x := range arg {
			cache = ping.SegmentKey(x)
			if err = gojo.ping.Handle(ctx, cache.TAG(ping.Anime), &result[i], func() error {
				result[i], err = q.GetAnimeTag(ctx, x)
				if err != nil {
					return err
				}

				return nil
			}); err != nil {
				ErrorSQL(err)
				return err
			}
		}

		return err
	})

	return result, err
}
