package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/ping"
)

func (gojo *SQLGojo) GetAllActorsTx(ctx context.Context, arg ListActorsParams) ([]Actor, error) {
	var err error
	var result []Actor
	var IDs []int64

	err = gojo.execTx(ctx, func(q *Queries) error {
		if err = gojo.ping.Handle(ctx, ping.CTM(ping.Anime, ping.ACT, arg.Limit, arg.Offset), IDs, func() error {
			IDs, err = q.ListActors(ctx, arg)
			if err != nil {
				return err
			}

			return nil
		}); err != nil {
			ErrorSQL(err)
			return err
		}

		if len(IDs) > 0 {
			var key ping.SegmentKey
			result = make([]Actor, len(IDs))
			for i, v := range IDs {
				key = ping.SegmentKey(v)
				if err = gojo.ping.Handle(ctx, key.ACT(), &result[i], func() error {
					result[i], err = q.GetActor(ctx, v)
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
