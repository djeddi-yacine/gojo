package db

import (
	"context"
)

func (gojo *SQLGojo) CreateGenresTx(ctx context.Context, arg []string) ([]Genre, error) {
	var result []Genre

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		for _, x := range arg {
			g, err := q.CreateGenre(ctx, x)
			if err != nil {
				return err
			}

			result = append(result, g)
		}

		return err
	})

	return result, err
}
