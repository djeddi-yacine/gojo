package db

import (
	"context"
	"fmt"

	"github.com/dj-yacine-flutter/gojo/ping"
)

func (gojo *SQLGojo) GetAllGenresTx(ctx context.Context, arg ListGenresParams) ([]Genre, error) {
	var err error
	var result []Genre
	var IDs []int32

	err = gojo.execTx(ctx, func(q *Queries) error {
		if err = gojo.ping.Handle(ctx, ping.CTM(ping.Anime, ping.GNR, fmt.Sprintf("%d-%d", arg.Limit, arg.Offset)), IDs, func() error {
			IDs, err = q.ListGenres(ctx, arg)
			if err != nil {
				return err
			}

			return nil
		}); err != nil {
			ErrorSQL(err)
			return err
		}

		if len(IDs) > 0 {
			result = make([]Genre, len(IDs))
			for i, v := range IDs {
				key := ping.SegmentKey(v)
				if err = gojo.ping.Handle(ctx, key.GNR(), &result[i], func() error {
					result[i], err = q.GetGenre(ctx, v)
					if err != nil {
						ErrorSQL(err)
						return err
					}

					return nil
				}); err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		return err
	})

	return result, err
}
