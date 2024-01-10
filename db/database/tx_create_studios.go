package db

import (
	"context"
)

func (gojo *SQLGojo) CreateStudiosTx(ctx context.Context, arg []string) ([]Studio, error) {
	var result []Studio

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		for _, x := range arg {
			s, err := q.CreateStudio(ctx, x)
			if err != nil {
				return err
			}
			result = append(result, s)
		}

		return err
	})

	return result, err
}
