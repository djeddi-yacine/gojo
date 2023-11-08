package db

import (
	"context"
	"errors"
)

type CreateAnimeSerieSeasonTxParams struct {
	Season      CreateAnimeSerieSeasonParams
	SeasonMetas []AnimeMetaTxParam
}

type CreateAnimeSerieSeasonTxResult struct {
	AnimeSerieSeason      AnimeSerieSeason
	AnimeSerieSeasonMetas []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeSerieSeasonTx(ctx context.Context, arg CreateAnimeSerieSeasonTxParams) (CreateAnimeSerieSeasonTxResult, error) {
	var result CreateAnimeSerieSeasonTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		_, err = q.GetAnimeSerie(ctx, arg.Season.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		season, err := q.CreateAnimeSerieSeason(ctx, arg.Season)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeSerieSeason = season

		if arg.SeasonMetas != nil {
			var metaArg CreateMetaParams
			var seasonMetaArg CreateAnimeSerieSeasonMetaParams
			result.AnimeSerieSeasonMetas = make([]AnimeMetaTxResult, len(arg.SeasonMetas))

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

				seasonMetaArg = CreateAnimeSerieSeasonMetaParams{
					SeasonID:   season.ID,
					LanguageID: m.LanguageID,
					MetaID:     meta.ID,
				}

				_, err = q.CreateAnimeSerieSeasonMeta(ctx, seasonMetaArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				l, err := q.GetLanguage(ctx, m.LanguageID)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeSerieSeasonMetas[i] = AnimeMetaTxResult{
					Meta:     meta,
					Language: l,
				}

			}

		} else {
			return errors.New("create one meta at least")
		}

		return err
	})

	return result, err
}
