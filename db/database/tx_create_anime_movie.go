package db

import (
	"context"

	"github.com/google/uuid"
)

type CreateAnimeMovieTxParams struct {
	CreateAnimeMovieParams    CreateAnimeMovieParams
	CreateAnimeResourceParams CreateAnimeResourceParams
	CreateAnimeLinkParams     CreateAnimeLinkParams
}

type CreateAnimeMovieTxResult struct {
	AnimeMovie    AnimeMovie
	AnimeResource AnimeResource
	AnimeLink     AnimeLink
}

func (gojo *SQLGojo) CreateAnimeMovieTx(ctx context.Context, arg CreateAnimeMovieTxParams) (CreateAnimeMovieTxResult, error) {
	var result CreateAnimeMovieTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		arg.CreateAnimeMovieParams.UniqueID = uuid.New()

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

		rarg := CreateAnimeMovieResourceParams{
			AnimeID:    anime.ID,
			ResourceID: resource.ID,
		}

		_, err = q.CreateAnimeMovieResource(ctx, rarg)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		link, err := q.CreateAnimeLink(ctx, arg.CreateAnimeLinkParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		larg := CreateAnimeMovieLinkParams{
			AnimeID: anime.ID,
			LinkID:  link.ID,
		}

		_, err = q.CreateAnimeMovieLink(ctx, larg)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeMovie = anime
		result.AnimeResource = resource
		result.AnimeLink = link

		return err
	})

	return result, err
}
