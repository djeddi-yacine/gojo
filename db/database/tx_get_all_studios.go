package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/ping"
)

func (gojo *SQLGojo) GetAllStudiosTx(ctx context.Context, arg ListStudiosParams) ([]Studio, error) {
	var err error
	var result []Studio
	var IDs []int32

	err = gojo.execTx(ctx, func(q *Queries) error {
		if err = gojo.ping.Handle(ctx, ping.CTM(ping.Anime, ping.STD, arg.Limit, arg.Offset), &IDs, func() error {
			IDs, err = q.ListStudios(ctx, arg)
			if err != nil {
				return err
			}

			return nil
		}); err != nil {
			ErrorSQL(err)
			return err
		}

		if len(IDs) > 0 {
			result = make([]Studio, len(IDs))
			for i, v := range IDs {
				if err = gojo.ping.Handle(ctx, ping.SegmentKey(v).STD(), &result[i], func() error {
					result[i], err = q.GetStudio(ctx, v)
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
