package db

import (
	"context"
)

type CreateAnimeSeasonInfoTxParams struct {
	SeasonID  int64
	GenreIDs  []int32
	StudioIDs []int32
}

type CreateAnimeSeasonInfoTxResult struct {
	AnimeSeason AnimeSeason
	Genres      []Genre
	Studios     []Studio
}

func (gojo *SQLGojo) CreateAnimeSeasonInfoTx(ctx context.Context, arg CreateAnimeSeasonInfoTxParams) (CreateAnimeSeasonInfoTxResult, error) {
	var result CreateAnimeSeasonInfoTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		result.AnimeSeason, err = q.GetAnimeSeason(ctx, arg.SeasonID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.GenreIDs != nil {
			var argGenre CreateAnimeSeasonGenreParams
			result.Genres = make([]Genre, len(arg.GenreIDs))
			for i, g := range arg.GenreIDs {
				argGenre = CreateAnimeSeasonGenreParams{
					SeasonID: result.AnimeSeason.ID,
					GenreID:  g,
				}

				_, err = q.CreateAnimeSeasonGenre(ctx, argGenre)
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
			var argStudio CreateAnimeSeasonStudioParams
			result.Studios = make([]Studio, len(arg.StudioIDs))
			for i, s := range arg.StudioIDs {
				argStudio = CreateAnimeSeasonStudioParams{
					SeasonID: result.AnimeSeason.ID,
					StudioID: s,
				}

				_, err = q.CreateAnimeSeasonStudio(ctx, argStudio)
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
