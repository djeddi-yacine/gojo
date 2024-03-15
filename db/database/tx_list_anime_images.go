package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/ping"
)

func (gojo *SQLGojo) ListAnimeImagesTx(ctx context.Context, arg []int64) ([]AnimeImage, error) {
	var err error
	var result []AnimeImage

	err = gojo.execTx(ctx, func(q *Queries) error {
		result = make([]AnimeImage, len(arg))

		for i, x := range arg {
			if err = gojo.ping.Handle(ctx, ping.SegmentKey(x).IMG(ping.Anime), &result[i], func() error {
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
