package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/ping"
)

func (gojo *SQLGojo) ListStudiosTx(ctx context.Context, arg []int32) ([]Studio, error) {
	var err error
	var result []Studio

	err = gojo.execTx(ctx, func(q *Queries) error {
		var cache ping.SegmentKey
		result = make([]Studio, len(arg))

		for i, x := range arg {
			cache = ping.SegmentKey(x)
			if err = gojo.ping.Handle(ctx, cache.STD(), &result[i], func() error {
				result[i], err = q.GetStudio(ctx, x)
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
