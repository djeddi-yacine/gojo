package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/ping"
)

func (gojo *SQLGojo) ListAnimeTrailersTx(ctx context.Context, arg []int64) ([]AnimeTrailer, error) {
	var err error
	var result []AnimeTrailer

	err = gojo.execTx(ctx, func(q *Queries) error {
		result = make([]AnimeTrailer, len(arg))

		for i, x := range arg {
			if err = gojo.ping.Handle(ctx, ping.SegmentKey(x).TRL(ping.Anime), &result[i], func() error {
				result[i], err = q.GetAnimeTrailer(ctx, x)
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
