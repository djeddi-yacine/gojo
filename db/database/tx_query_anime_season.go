package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
)

type QueryAnimeSeasonTxParams struct {
	Query  string
	Limit  int32
	Offset int32
}

type QueryAnimeSeasonTxResult struct {
	AnimeSeasons []AnimeSeason
}

func (gojo *SQLGojo) QueryAnimeSeasonTx(ctx context.Context, arg QueryAnimeSeasonTxParams) (QueryAnimeSeasonTxResult, error) {
	var result QueryAnimeSeasonTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		var animeSeasonIDs []int64

		animeOfficialResults, err := q.QueryAnimeSeasonOfficialTitles(ctx, QueryAnimeSeasonOfficialTitlesParams{
			Column1: arg.Query,
			Limit:   arg.Limit,
			Offset:  arg.Offset,
		})
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if len(animeOfficialResults) > 0 {
			animeSeasonIDs = append(animeSeasonIDs, animeOfficialResults...)
		}

		animeShortResults, err := q.QueryAnimeSeasonShortTitles(ctx, QueryAnimeSeasonShortTitlesParams{
			Column1: arg.Query,
			Limit:   arg.Limit,
			Offset:  arg.Offset,
		})
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if len(animeShortResults) > 0 {
			animeSeasonIDs = append(animeSeasonIDs, animeShortResults...)
		}

		animeOtherResults, err := q.QueryAnimeSeasonOtherTitles(ctx, QueryAnimeSeasonOtherTitlesParams{
			Column1: arg.Query,
			Limit:   arg.Limit,
			Offset:  arg.Offset,
		})
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if len(animeOtherResults) > 0 {
			animeSeasonIDs = append(animeSeasonIDs, animeOtherResults...)
		}

		animeTranslationResults, err := q.QueryAnimeSeasonTranslationTitles(ctx, QueryAnimeSeasonTranslationTitlesParams{
			Column1: arg.Query,
			Limit:   arg.Limit,
			Offset:  arg.Offset,
		})
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if len(animeTranslationResults) > 0 {
			animeSeasonIDs = append(animeSeasonIDs, animeTranslationResults...)
		}

		IDs := utils.RemoveDuplicatesInt64(animeSeasonIDs)

		if len(IDs) > 0 {
			result.AnimeSeasons = make([]AnimeSeason, len(IDs))

			var animeSeason AnimeSeason

			for i, id := range IDs {
				animeSeason, err = q.GetAnimeSeason(ctx, id)
				if err != nil {
					if ErrorDB(err).Code != pgerrcode.CaseNotFound {
						ErrorSQL(err)
						return err
					}
					continue
				}
				result.AnimeSeasons[i] = animeSeason
			}
		}

		return err
	})

	return result, err
}
