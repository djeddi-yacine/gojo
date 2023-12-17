package db

import (
	"context"
	"errors"
)

type CreateAnimeSeasonMetasTxParams struct {
	SeasonID    int64
	SeasonMetas []AnimeMetaTxParam
}

type CreateAnimeSeasonMetasTxResult struct {
	AnimeSeason      AnimeSerieSeason
	AnimeSeasonMetas []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeSeasonMetasTx(ctx context.Context, arg CreateAnimeSeasonMetasTxParams) (CreateAnimeSeasonMetasTxResult, error) {
	var result CreateAnimeSeasonMetasTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		season, err := q.GetAnimeSeason(ctx, arg.SeasonID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeSeason = season

		if arg.SeasonMetas != nil {
			var metaArg CreateMetaParams
			var seasonMetaArg CreateAnimeSeasonMetaParams
			result.AnimeSeasonMetas = make([]AnimeMetaTxResult, len(arg.SeasonMetas))

			for i, m := range arg.SeasonMetas {
				metaArg = CreateMetaParams{
					Title:    m.Title,
					Overview: m.Overview,
				}

				meta, err := q.CreateMeta(ctx, metaArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				seasonMetaArg = CreateAnimeSeasonMetaParams{
					SeasonID:   season.ID,
					LanguageID: m.LanguageID,
					MetaID:     meta.ID,
				}

				_, err = q.CreateAnimeSeasonMeta(ctx, seasonMetaArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeSeasonMetas[i] = AnimeMetaTxResult{
					Meta:       meta,
					LanguageID: m.LanguageID,
				}

				_, err = q.CreateAnimeSeasonTranslationTitle(ctx, CreateAnimeSeasonTranslationTitleParams{
					SeasonID:  arg.SeasonID,
					TitleText: meta.Title,
				})
				if err != nil {
					ErrorSQL(err)
					return err
				}
			}

		} else {
			return errors.New("create one meta at least")
		}

		return err
	})

	return result, err
}
