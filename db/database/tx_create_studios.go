package db

import (
	"context"
)

type CreateStudiosTxParams struct {
	Names []string
}

type CreateStudiosTxResult struct {
	Studios []Studio
}

func (gojo *SQLGojo) CreateStudiosTx(ctx context.Context, arg CreateStudiosTxParams) (CreateStudiosTxResult, error) {
	var result CreateStudiosTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		for _, name := range arg.Names {
			s, err := q.CreateStudio(ctx, name)
			if err != nil {
				return err
			}
			result.Studios = append(result.Studios, s)
		}

		return err
	})

	return result, err
}
