package db

import (
	"context"
)

type CreateAnimeSerieTxParams struct {
	CreateAnimeSerieParams    CreateAnimeSerieParams
	CreateAnimeResourceParams CreateAnimeResourceParams
}

type CreateAnimeSerieTxResult struct {
	AnimeSerie AnimeSerie
	Resource   AnimeResource
}

func (gojo *SQLGojo) CreateAnimeSerieTx(ctx context.Context, arg CreateAnimeSerieTxParams) (CreateAnimeSerieTxResult, error) {
	var result CreateAnimeSerieTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.CreateAnimeSerie(ctx, arg.CreateAnimeSerieParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		resource, err := q.CreateAnimeResource(ctx, arg.CreateAnimeResourceParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		arg := CreateAnimeSerieResourceParams{
			AnimeID:    anime.ID,
			ResourceID: resource.ID,
		}

		_, err = q.CreateAnimeSerieResource(ctx, arg)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeSerie = anime
		result.Resource = resource

		return err
	})

	return result, err
}
