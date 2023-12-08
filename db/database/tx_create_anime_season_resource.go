package db

import (
	"context"
)

type CreateAnimeSeasonResourceTxParams struct {
	SeasonID                  int64
	CreateAnimeResourceParams CreateAnimeResourceParams
}

type CreateAnimeSeasonResourceTxResult struct {
	AnimeSeason   AnimeSerieSeason
	AnimeResource AnimeResource
}

func (gojo *SQLGojo) CreateAnimeSeasonResourceTx(ctx context.Context, arg CreateAnimeSeasonResourceTxParams) (CreateAnimeSeasonResourceTxResult, error) {
	var result CreateAnimeSeasonResourceTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.GetAnimeSerieSeason(ctx, arg.SeasonID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		resource, err := q.CreateAnimeResource(ctx, arg.CreateAnimeResourceParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		arg := CreateAnimeSeasonResourceParams{
			SeasonID:   arg.SeasonID,
			ResourceID: resource.ID,
		}

		_, err = q.CreateAnimeSeasonResource(ctx, arg)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeResource = resource
		result.AnimeSeason = anime

		return err
	})

	return result, err
}
