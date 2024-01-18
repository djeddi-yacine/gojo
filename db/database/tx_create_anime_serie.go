package db

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type CreateAnimeSerieTxParams struct {
	CreateAnimeSerieParams CreateAnimeSerieParams
	CreateAnimeLinkParams  CreateAnimeLinkParams
	CreateAnimeMetasParams []AnimeMetaTxParam
}

type CreateAnimeSerieTxResult struct {
	AnimeSerie AnimeSerie
	AnimeLink  AnimeLink
	AnimeMetas []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeSerieTx(ctx context.Context, arg CreateAnimeSerieTxParams) (CreateAnimeSerieTxResult, error) {
	var result CreateAnimeSerieTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		arg.CreateAnimeSerieParams.UniqueID = uuid.New()

		anime, err := q.CreateAnimeSerie(ctx, arg.CreateAnimeSerieParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		link, err := q.CreateAnimeLink(ctx, arg.CreateAnimeLinkParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		larg := CreateAnimeSerieLinkParams{
			AnimeID: anime.ID,
			LinkID:  link.ID,
		}

		_, err = q.CreateAnimeSerieLink(ctx, larg)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.CreateAnimeMetasParams != nil {
			result.AnimeMetas = make([]AnimeMetaTxResult, len(arg.CreateAnimeMetasParams))

			for i, m := range arg.CreateAnimeMetasParams {
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
					AnimeID:    anime.ID,
					LanguageID: lang.ID,
					MetaID:     meta.ID,
				}
				animeMeta, err := q.CreateAnimeSerieMeta(ctx, arg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeMetas[i].Meta, err = q.GetMeta(ctx, animeMeta.MetaID)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeMetas[i].LanguageID = animeMeta.LanguageID
			}
		} else {
			return errors.New("create one meta at least")
		}

		result.AnimeSerie = anime
		result.AnimeLink = link

		return err
	})

	return result, err
}
