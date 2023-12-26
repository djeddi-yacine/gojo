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
	AnimeSeasonMetas []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeSeasonMetasTx(ctx context.Context, arg CreateAnimeSeasonMetasTxParams) (CreateAnimeSeasonMetasTxResult, error) {
	var result CreateAnimeSeasonMetasTxResult
	var err error

	err = gojo.execTx(ctx, func(q *Queries) error {
		_, err = q.GetAnimeSeason(ctx, arg.SeasonID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.SeasonMetas != nil {
			result.AnimeSeasonMetas = make([]AnimeMetaTxResult, len(arg.SeasonMetas))

			for i, m := range arg.SeasonMetas {
				lang, err := q.GetLanguage(ctx, m.LanguageID)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				meta, err := q.CreateMeta(ctx, m.CreateMetaParams)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				seasonMetaArg := CreateAnimeSeasonMetaParams{
					SeasonID:   arg.SeasonID,
					LanguageID: lang.ID,
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
