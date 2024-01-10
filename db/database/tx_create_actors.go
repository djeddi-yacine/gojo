package db

import (
	"context"
)

type CreateActorsTxParams struct {
	Actors []CreateActorParams
}

type CreateActorsTxResult struct {
	Actors []Actor
}

func (gojo *SQLGojo) CreateActorsTx(ctx context.Context, arg CreateActorsTxParams) (CreateActorsTxResult, error) {
	var result CreateActorsTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		for _, x := range arg.Actors {
			z, err := q.CreateActor(ctx, x)
			if err != nil {
				return err
			}

			result.Actors = append(result.Actors, z)
		}

		return err
	})

	return result, err
}
