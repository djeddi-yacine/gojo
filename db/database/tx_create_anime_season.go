package db

import (
	"context"
	"errors"
)

type CreateAnimeSeasonTxParams struct {
	Season      CreateAnimeSeasonParams
	SeasonMetas []AnimeMetaTxParam
}

type CreateAnimeSeasonTxResult struct {
	AnimeSeason      AnimeSeason
	AnimeSeasonMetas []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeSeasonTx(ctx context.Context, arg CreateAnimeSeasonTxParams) (CreateAnimeSeasonTxResult, error) {
	var result CreateAnimeSeasonTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		season, err := q.CreateAnimeSeason(ctx, arg.Season)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		_, err = q.CreateAnimeSerieSeason(ctx, CreateAnimeSerieSeasonParams{
			AnimeID:  season.AnimeID,
			SeasonID: season.ID,
		})

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

			}

		} else {
			return errors.New("create one meta at least")
		}

		return err
	})

	return result, err
}
