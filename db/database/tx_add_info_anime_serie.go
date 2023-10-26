package db

import (
	"context"
)

type AddInfoAnimeSerieTxParams struct {
	AnimeID   int64
	GenreIDs  []int32
	StudioIDs []int32
}

type AddInfoAnimeSerieTxResult struct {
	AnimeSerie
	AnimeSerieGenres  []AnimeSerieGenre
	AnimeSerieStudios []AnimeSerieStudio
}

func (gojo *SQLGojo) AddInfoAnimeSerieTx(ctx context.Context, arg AddInfoAnimeSerieTxParams) (AddInfoAnimeSerieTxResult, error) {
	var result AddInfoAnimeSerieTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		result.AnimeSerie, err = q.GetAnimeSerie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.GenreIDs != nil {
			var argGenre CreateAnimeSerieGenreParams
			result.AnimeSerieGenres = make([]AnimeSerieGenre, len(arg.GenreIDs))
			for i, g := range arg.GenreIDs {
				argGenre = CreateAnimeSerieGenreParams{
					AnimeID: result.AnimeSerie.ID,
					GenreID: g,
				}
				ag, err := q.CreateAnimeSerieGenre(ctx, argGenre)
				if err != nil {
					ErrorSQL(err)
					return err
				}
				result.AnimeSerieGenres[i] = ag
			}
		}

		if arg.StudioIDs != nil {
			var argStudio CreateAnimeSerieStudioParams
			result.AnimeSerieStudios = make([]AnimeSerieStudio, len(arg.StudioIDs))
			for i, s := range arg.StudioIDs {
				argStudio = CreateAnimeSerieStudioParams{
					AnimeID:  result.AnimeSerie.ID,
					StudioID: s,
				}

				as, err := q.CreateAnimeSerieStudio(ctx, argStudio)
				if err != nil {
					ErrorSQL(err)
					return err
				}
				result.AnimeSerieStudios[i] = as
			}
		}

		return err
	})

	return result, err
}
