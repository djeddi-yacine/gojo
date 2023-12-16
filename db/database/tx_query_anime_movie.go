package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
)

type QueryAnimeMovieTxParams struct {
	Query string
}

type QueryAnimeMovieTxResult struct {
	AnimeMovies []AnimeMovie
}

func (gojo *SQLGojo) QueryAnimeMovieTx(ctx context.Context, arg QueryAnimeMovieTxParams) (QueryAnimeMovieTxResult, error) {
	var result QueryAnimeMovieTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		var animeMovieIDs []int64

		animeOfficialResults, err := q.QueryAnimeMovieOfficialTitles(ctx, arg.Query)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if len(animeOfficialResults) > 0 {
			animeMovieIDs = append(animeMovieIDs, animeOfficialResults...)
		}

		animeShortResults, err := q.QueryAnimeMovieShortTitles(ctx, arg.Query)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if len(animeShortResults) > 0 {
			animeMovieIDs = append(animeMovieIDs, animeShortResults...)
		}

		animeOtherResults, err := q.QueryAnimeMovieOtherTitles(ctx, arg.Query)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if len(animeOtherResults) > 0 {
			animeMovieIDs = append(animeMovieIDs, animeOtherResults...)
		}

		animeTranslationResults, err := q.QueryAnimeMovieTranslationTitles(ctx, arg.Query)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if len(animeTranslationResults) > 0 {
			animeMovieIDs = append(animeMovieIDs, animeTranslationResults...)
		}

		IDs := utils.RemoveDuplicatesINT64(animeMovieIDs)

		if len(IDs) > 0 {
			result.AnimeMovies = make([]AnimeMovie, len(IDs))

			var animeMovie AnimeMovie

			for i, id := range IDs {
				animeMovie, err = q.GetAnimeMovie(ctx, id)
				if err != nil {
					if ErrorDB(err).Code != pgerrcode.CaseNotFound {
						ErrorSQL(err)
						return err
					}
					continue
				}
				result.AnimeMovies[i] = animeMovie
			}
		}

		return err
	})

	return result, err
}
