package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/ping"
)

func (gojo *SQLGojo) ListGenresTx(ctx context.Context, arg []int32) ([]Genre, error) {
	var err error
	var result []Genre

	err = gojo.execTx(ctx, func(q *Queries) error {
		var cache ping.SegmentKey
		result = make([]Genre, len(arg))

		for i, x := range arg {
			cache = ping.SegmentKey(x)
			if err = gojo.ping.Handle(ctx, cache.GNR(), &result[i], func() error {
				result[i], err = q.GetGenre(ctx, x)
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
