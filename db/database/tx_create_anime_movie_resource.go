package db

import (
	"context"
)

type CreateAnimeMovieResourceTxParams struct {
	AnimeID                   int64
	CreateAnimeResourceParams CreateAnimeResourceParams
}

type CreateAnimeMovieResourceTxResult struct {
	AnimeMovie    AnimeMovie
	AnimeResource AnimeResource
}

func (gojo *SQLGojo) CreateAnimeMovieResourceTx(ctx context.Context, arg CreateAnimeMovieResourceTxParams) (CreateAnimeMovieResourceTxResult, error) {
	var result CreateAnimeMovieResourceTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.GetAnimeMovie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		resource, err := q.CreateAnimeResource(ctx, arg.CreateAnimeResourceParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		arg := CreateAnimeMovieResourceParams{
			AnimeID:    arg.AnimeID,
			ResourceID: resource.ID,
		}

		_, err = q.CreateAnimeMovieResource(ctx, arg)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeResource = resource
		result.AnimeMovie = anime

		return err
	})

	return result, err
}
