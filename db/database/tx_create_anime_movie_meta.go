package db

import (
	"context"
)

type CreateAnimeMovieMetaTxParams struct {
	AnimeID    int64
	LanguageID int32
	CreateMetaParams
}

type CreateAnimeMovieMetaTxResult struct {
	Language Language
	Meta Meta
}

func (gojo *SQLGojo) CreateAnimeMovieMetaTx(ctx context.Context, arg CreateAnimeMovieMetaTxParams) (CreateAnimeMovieMetaTxResult, error) {
	var result CreateAnimeMovieMetaTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		meta, err := q.CreateMeta(ctx, arg.CreateMetaParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		arg := CreateAnimeMetaParams{
			AnimeID:    arg.AnimeID,
			LanguageID: arg.LanguageID,
			MetaID:     meta.ID,
		}
		animeMeta, err := q.CreateAnimeMeta(ctx, arg)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.Meta, err = q.GetMeta(ctx, animeMeta.MetaID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.Language, err = q.GetLanguage(ctx, animeMeta.LanguageID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		return err
	})

	return result, err
}
