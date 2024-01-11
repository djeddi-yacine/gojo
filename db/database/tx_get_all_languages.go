package db

import (
	"context"
	"fmt"

	"github.com/dj-yacine-flutter/gojo/ping"
)

func (gojo *SQLGojo) GetAllLanguagesTx(ctx context.Context, arg ListLanguagesParams) ([]Language, error) {
	var err error
	var result []Language
	var IDs []int32

	err = gojo.execTx(ctx, func(q *Queries) error {
		if err = gojo.ping.Handle(ctx, ping.CTM(ping.Anime, ping.LNG, fmt.Sprintf("%d-%d", arg.Limit, arg.Offset)), IDs, func() error {
			IDs, err = q.ListLanguages(ctx, arg)
			if err != nil {
				return err
			}

			return nil
		}); err != nil {
			ErrorSQL(err)
			return err
		}

		if len(IDs) > 0 {
			result = make([]Language, len(IDs))
			for i, v := range IDs {
				key := ping.SegmentKey(v)
				if err = gojo.ping.Handle(ctx, key.LNG(), &result[i], func() error {
					result[i], err = q.GetLanguage(ctx, v)
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
