package db

import (
	"context"
)

type AddAnimeMovieInfoTxParams struct {
	AnimeID   int64
	GenreIDs  []int32
	StudioIDs []int32
}

type AddAnimeMovieInfoTxResult struct {
	AnimeMovie
	Genres  []Genre
	Studios []Studio
}

func (gojo *SQLGojo) AddAnimeMovieInfoTx(ctx context.Context, arg AddAnimeMovieInfoTxParams) (AddAnimeMovieInfoTxResult, error) {
	var result AddAnimeMovieInfoTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		result.AnimeMovie, err = q.GetAnimeMovie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.GenreIDs != nil {
			var argGenre CreateAnimeMovieGenreParams
			result.Genres = make([]Genre, len(arg.GenreIDs))
			for i, g := range arg.GenreIDs {
				argGenre = CreateAnimeMovieGenreParams{
					AnimeID: result.AnimeMovie.ID,
					GenreID: g,
				}

				_, err = q.CreateAnimeMovieGenre(ctx, argGenre)
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
			var argStudio CreateAnimeMovieStudioParams
			result.Studios = make([]Studio, len(arg.StudioIDs))
			for i, s := range arg.StudioIDs {
				argStudio = CreateAnimeMovieStudioParams{
					AnimeID:  result.AnimeMovie.ID,
					StudioID: s,
				}

				_, err = q.CreateAnimeMovieStudio(ctx, argStudio)
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
