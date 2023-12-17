package db

import (
	"context"

	"github.com/dj-yacine-flutter/gojo/utils"
	"github.com/jackc/pgerrcode"
)

type QueryAnimeSeasonTxParams struct {
	Query string
}

type QueryAnimeSeasonTxResult struct {
	AnimeSeasons []AnimeSerieSeason
}

func (gojo *SQLGojo) QueryAnimeSeasonTx(ctx context.Context, arg QueryAnimeSeasonTxParams) (QueryAnimeSeasonTxResult, error) {
	var result QueryAnimeSeasonTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		var animeSeasonIDs []int64

		animeOfficialResults, err := q.QueryAnimeSeasonOfficialTitles(ctx, arg.Query)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if len(animeOfficialResults) > 0 {
			animeSeasonIDs = append(animeSeasonIDs, animeOfficialResults...)
		}

		animeShortResults, err := q.QueryAnimeSeasonShortTitles(ctx, arg.Query)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if len(animeShortResults) > 0 {
			animeSeasonIDs = append(animeSeasonIDs, animeShortResults...)
		}

		animeOtherResults, err := q.QueryAnimeSeasonOtherTitles(ctx, arg.Query)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if len(animeOtherResults) > 0 {
			animeSeasonIDs = append(animeSeasonIDs, animeOtherResults...)
		}

		animeTranslationResults, err := q.QueryAnimeSeasonTranslationTitles(ctx, arg.Query)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if len(animeTranslationResults) > 0 {
			animeSeasonIDs = append(animeSeasonIDs, animeTranslationResults...)
		}

		IDs := utils.RemoveDuplicatesINT64(animeSeasonIDs)

		if len(IDs) > 0 {
			result.AnimeSeasons = make([]AnimeSerieSeason, len(IDs))

			var animeSeason AnimeSerieSeason

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
