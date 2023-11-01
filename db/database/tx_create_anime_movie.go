package db

import (
	"context"
)

type CreateAnimeMovieTxParams struct {
	CreateAnimeMovieParams    CreateAnimeMovieParams
	CreateAnimeResourceParams CreateAnimeResourceParams
}

type CreateAnimeMovieTxResult struct {
	AnimeMovie AnimeMovie
	Resource   AnimeResource
}

func (gojo *SQLGojo) CreateAnimeMovieTx(ctx context.Context, arg CreateAnimeMovieTxParams) (CreateAnimeMovieTxResult, error) {
	var result CreateAnimeMovieTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.CreateAnimeMovie(ctx, arg.CreateAnimeMovieParams)
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
			AnimeID:    anime.ID,
			ResourceID: resource.ID,
		}

		_, err = q.CreateAnimeMovieResource(ctx, arg)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeMovie = anime
		result.Resource = resource

		return err
	})

	return result, err
}
