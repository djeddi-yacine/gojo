package db

import (
	"context"
)

func (gojo *SQLGojo) CreateActorsTx(ctx context.Context, arg []CreateActorParams) ([]Actor, error) {
	var result []Actor

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		for _, x := range arg {
			z, err := q.CreateActor(ctx, x)
			if err != nil {
				ErrorSQL(err)
				return err
			}

			result = append(result, z)
		}

		return err
	})

	return result, err
}
