package db

import (
	"context"
)

type CreateGenresTxParams struct {
	Names []string
}

type CreateGenresTxResult struct {
	Genres []Genre
}

func (gojo *SQLGojo) CreateGenresTx(ctx context.Context, arg CreateGenresTxParams) (CreateGenresTxResult, error) {
	var result CreateGenresTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		for _, name := range arg.Names {
			g, err := q.CreateGenre(ctx, name)
			if err != nil {
				return err
			}
			result.Genres = append(result.Genres, g)
		}

		return err
	})

	return result, err
}
