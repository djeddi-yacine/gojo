package db

import (
	"context"
	"errors"
)

type CreateAnimeSerieMetasTxParams struct {
	AnimeID         int64
	AnimeSerieMetas []AnimeMetaTxParam
}

type CreateAnimeSerieMetasTxResult struct {
	AnimeSerieMetas []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeSerieMetasTx(ctx context.Context, arg CreateAnimeSerieMetasTxParams) (CreateAnimeSerieMetasTxResult, error) {
	var result CreateAnimeSerieMetasTxResult
	var err error

	err = gojo.execTx(ctx, func(q *Queries) error {
		_, err = q.GetAnimeSerie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.AnimeSerieMetas != nil {
			result.AnimeSerieMetas = make([]AnimeMetaTxResult, len(arg.AnimeSerieMetas))

			for i, m := range arg.AnimeSerieMetas {
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

				arg := CreateAnimeSerieMetaParams{
					AnimeID:    arg.AnimeID,
					LanguageID: lang.ID,
					MetaID:     meta.ID,
				}
				animeMeta, err := q.CreateAnimeSerieMeta(ctx, arg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeSerieMetas[i].Meta, err = q.GetMeta(ctx, animeMeta.MetaID)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeSerieMetas[i].LanguageID = animeMeta.LanguageID
			}
		} else {
			return errors.New("create one meta at least")
		}

		return err
	})

	return result, err
}
