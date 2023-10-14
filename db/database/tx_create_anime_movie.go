package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type CreateAnimeMovieTxParams struct {
	CreateAnimeMovieParams
	GenreIDs  []int32
	StudioIDs []int32
}

type CreateAnimeMovieTxResult struct {
	AnimeMovie AnimeMovie
}

func (gojo *SQLGojo) CreateAnimeMovieTx(ctx context.Context, arg CreateAnimeMovieTxParams) (CreateAnimeMovieTxResult, error) {
	var result CreateAnimeMovieTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		result.AnimeMovie, err = q.CreateAnimeMovie(ctx, arg.CreateAnimeMovieParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.GenreIDs != nil {
			var argGenre CreateAnimeGenreParams
			for _, g := range arg.GenreIDs {
				argGenre = CreateAnimeGenreParams{
					AnimeID: result.AnimeMovie.ID,
					GenreID: pgtype.Int4{
						Int32: g,
						Valid: true,
					},
				}
				_, err = q.CreateAnimeGenre(ctx, argGenre)
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		if arg.StudioIDs != nil {
			var argStudio CreateAnimeStudioParams
			for _, s := range arg.StudioIDs {
				argStudio = CreateAnimeStudioParams{
					AnimeID: result.AnimeMovie.ID,
					StudioID: pgtype.Int4{
						Int32: s,
						Valid: true,
					},
				}

				_, err = q.CreateAnimeStudio(ctx, argStudio)
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}
		}

		return err
	})

	return result, err
}
