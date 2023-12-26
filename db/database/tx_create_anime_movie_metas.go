package db

import (
	"context"
	"errors"
)

type CreateAnimeMovieMetasTxParams struct {
	AnimeID         int64
	AnimeMovieMetas []AnimeMetaTxParam
}

type CreateAnimeMovieMetasTxResult struct {
	AnimeMovieMetas []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeMovieMetasTx(ctx context.Context, arg CreateAnimeMovieMetasTxParams) (CreateAnimeMovieMetasTxResult, error) {
	var result CreateAnimeMovieMetasTxResult
	var err error

	err = gojo.execTx(ctx, func(q *Queries) error {
		_, err = q.GetAnimeMovie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		if arg.AnimeMovieMetas != nil {
			result.AnimeMovieMetas = make([]AnimeMetaTxResult, len(arg.AnimeMovieMetas))

			for i, m := range arg.AnimeMovieMetas {
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

				arg := CreateAnimeMovieMetaParams{
					AnimeID:    arg.AnimeID,
					LanguageID: lang.ID,
					MetaID:     meta.ID,
				}
				animeMeta, err := q.CreateAnimeMovieMeta(ctx, arg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeMovieMetas[i].Meta, err = q.GetMeta(ctx, animeMeta.MetaID)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeMovieMetas[i].LanguageID = animeMeta.LanguageID

				_, err = q.CreateAnimeMovieTranslationTitle(ctx, CreateAnimeMovieTranslationTitleParams{
					AnimeID:   arg.AnimeID,
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
