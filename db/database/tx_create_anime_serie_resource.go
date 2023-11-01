package db

import (
	"context"
)

type CreateAnimeSerieResourceTxParams struct {
	AnimeID                   int64
	CreateAnimeResourceParams CreateAnimeResourceParams
}

type CreateAnimeSerieResourceTxResult struct {
	AnimeSerie    AnimeSerie
	AnimeResource AnimeResource
}

func (gojo *SQLGojo) CreateAnimeSerieResourceTx(ctx context.Context, arg CreateAnimeSerieResourceTxParams) (CreateAnimeSerieResourceTxResult, error) {
	var result CreateAnimeSerieResourceTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.GetAnimeSerie(ctx, arg.AnimeID)
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
			AnimeID:    arg.AnimeID,
			ResourceID: resource.ID,
		}

		_, err = q.CreateAnimeSerieResource(ctx, arg)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeResource = resource
		result.AnimeSerie = anime

		return err
	})

	return result, err
}
