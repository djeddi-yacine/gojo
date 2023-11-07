package db

import (
	"context"
)

type AddAnimeSerieInfoTxParams struct {
	AnimeID   int64
	GenreIDs  []int32
	StudioIDs []int32
}

type AddAnimeSerieInfoTxResult struct {
	AnimeSerie
	Genres  []Genre
	Studios []Studio
}

func (gojo *SQLGojo) AddAnimeSerieInfoTx(ctx context.Context, arg AddAnimeSerieInfoTxParams) (AddAnimeSerieInfoTxResult, error) {
	var result AddAnimeSerieInfoTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		result.AnimeSerie, err = q.GetAnimeSerie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.GenreIDs != nil {
			var argGenre CreateAnimeSerieGenreParams
			result.Genres = make([]Genre, len(arg.GenreIDs))
			for i, g := range arg.GenreIDs {
				argGenre = CreateAnimeSerieGenreParams{
					AnimeID: result.AnimeSerie.ID,
					GenreID: g,
				}

				_, err = q.CreateAnimeSerieGenre(ctx, argGenre)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				ng, err := q.GetGenre(ctx, g)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.Genres[i] = ng
			}
		}

		if arg.StudioIDs != nil {
			var argStudio CreateAnimeSerieStudioParams
			result.Studios = make([]Studio, len(arg.StudioIDs))
			for i, s := range arg.StudioIDs {
				argStudio = CreateAnimeSerieStudioParams{
					AnimeID:  result.AnimeSerie.ID,
					StudioID: s,
				}

				_, err = q.CreateAnimeSerieStudio(ctx, argStudio)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				ns, err := q.GetStudio(ctx, s)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.Studios[i] = ns
			}
		}

		return err
	})

	return result, err
}
