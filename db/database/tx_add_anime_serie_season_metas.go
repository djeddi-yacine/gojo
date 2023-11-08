package db

import (
	"context"
	"errors"
	"fmt"
	"math"
)

type AddAnimeSerieSeasonMetasTxParams struct {
	SeasonID    int64
	SeasonMetas []AnimeMetaTxParam
}

type AddAnimeSerieSeasonMetasTxResult struct {
	AnimeSerieSeason      AnimeSerieSeason
	AnimeSerieSeasonMetas []AnimeMetaTxResult
}

func (gojo *SQLGojo) AddAnimeSerieSeasonMetasTx(ctx context.Context, arg AddAnimeSerieSeasonMetasTxParams) (AddAnimeSerieSeasonMetasTxResult, error) {
	var result AddAnimeSerieSeasonMetasTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		season, err := q.GetAnimeSerieSeason(ctx, arg.SeasonID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		languages, err := q.ListLanguages(ctx, ListLanguagesParams{
			Limit:  math.MaxInt32,
			Offset: 0,
		})
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
				if checkLanguage(languages, m.LanguageID) {
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

				} else {
					return fmt.Errorf("there is no language with ID : %d", m.LanguageID)
				}
			}

		} else {
			return errors.New("create one meta at least")
		}

		return err
	})

	return result, err
}
