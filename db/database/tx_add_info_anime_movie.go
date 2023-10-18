package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type AddInfoAnimeMovieTxParams struct {
	AnimeID   int64
	GenreIDs  []int32
	StudioIDs []int32
}

type AddInfoAnimeMovieTxResult struct {
	AnimeMovie
	AnimeGenres  []AnimeGenre
	AnimeStudios []AnimeStudio
}

func (gojo *SQLGojo) AddInfoAnimeMovieTx(ctx context.Context, arg AddInfoAnimeMovieTxParams) (AddInfoAnimeMovieTxResult, error) {
	var result AddInfoAnimeMovieTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		result.AnimeMovie, err = q.GetAnimeMovie(ctx, arg.AnimeID)
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
				ag, err := q.CreateAnimeGenre(ctx, argGenre)
				if err != nil {
					ErrorSQL(err)
					return err
				}
				result.AnimeGenres = append(result.AnimeGenres, ag)
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

				as, err := q.CreateAnimeStudio(ctx, argStudio)
				if err != nil {
					ErrorSQL(err)
					return err
				}
				result.AnimeStudios = append(result.AnimeStudios, as)
			}
		}

		return err
	})

	return result, err
}
