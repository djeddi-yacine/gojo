package db

import (
	"context"
)

func (gojo *SQLGojo) CreateLanguagesTx(ctx context.Context, arg []CreateLanguageParams) ([]Language, error) {
	var result []Language

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		for _, x := range arg {
			l, err := q.CreateLanguage(ctx, x)
			if err != nil {
				return err
			}

			result = append(result, l)
		}

		return err
	})

	return result, err
}
