package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/ping"
)

func (gojo *SQLGojo) ListAnimeImagesTx(ctx context.Context, arg []int64) ([]AnimeImage, error) {
	var err error
	var result []AnimeImage

	err = gojo.execTx(ctx, func(q *Queries) error {
		var cache ping.SegmentKey
		result = make([]AnimeImage, len(arg))

		for i, x := range arg {
			cache = ping.SegmentKey(x)
			if err = gojo.ping.Handle(ctx, cache.IMG(ping.Anime), &result[i], func() error {
				result[i], err = q.GetAnimeImage(ctx, x)
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
