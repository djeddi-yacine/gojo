package db

import (
	"context"
)

type AddInfoAnimeMovieTxParams struct {
	AnimeID   int64
	GenreIDs  []int32
	StudioIDs []int32
}

type AddInfoAnimeMovieTxResult struct {
	AnimeMovie
	AnimeMovieGenres  []AnimeMovieGenre
	AnimeMovieStudios []AnimeMovieStudio
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
			var argGenre CreateAnimeMovieGenreParams
			result.AnimeMovieGenres = make([]AnimeMovieGenre, len(arg.GenreIDs))
			for i, g := range arg.GenreIDs {
				argGenre = CreateAnimeMovieGenreParams{
					AnimeID: result.AnimeMovie.ID,
					GenreID: g,
				}
				ag, err := q.CreateAnimeMovieGenre(ctx, argGenre)
				if err != nil {
					ErrorSQL(err)
					return err
				}
				result.AnimeMovieGenres[i] = ag
			}
		}

		if arg.StudioIDs != nil {
			var argStudio CreateAnimeMovieStudioParams
			result.AnimeMovieStudios = make([]AnimeMovieStudio, len(arg.StudioIDs))
			for i, s := range arg.StudioIDs {
				argStudio = CreateAnimeMovieStudioParams{
					AnimeID:  result.AnimeMovie.ID,
					StudioID: s,
				}

				as, err := q.CreateAnimeMovieStudio(ctx, argStudio)
				if err != nil {
					ErrorSQL(err)
					return err
				}
				result.AnimeMovieStudios[i] = as
			}
		}

		return err
	})

	return result, err
}
