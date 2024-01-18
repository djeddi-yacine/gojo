package db

import (
	"context"
	"errors"

	"github.com/google/uuid"
)

type CreateAnimeMovieTxParams struct {
	CreateAnimeMovieParams    CreateAnimeMovieParams
	CreateAnimeResourceParams CreateAnimeResourceParams
	CreateAnimeLinkParams     CreateAnimeLinkParams
	CreateAnimeMetasParams    []AnimeMetaTxParam
}

type CreateAnimeMovieTxResult struct {
	AnimeMovie    AnimeMovie
	AnimeResource AnimeResource
	AnimeLink     AnimeLink
	AnimeMetas    []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeMovieTx(ctx context.Context, arg CreateAnimeMovieTxParams) (CreateAnimeMovieTxResult, error) {
	var result CreateAnimeMovieTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		arg.CreateAnimeMovieParams.UniqueID = uuid.New()

		anime, err := q.CreateAnimeMovie(ctx, arg.CreateAnimeMovieParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		resource, err := q.CreateAnimeResource(ctx, arg.CreateAnimeResourceParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		rarg := CreateAnimeMovieResourceParams{
			AnimeID:    anime.ID,
			ResourceID: resource.ID,
		}

		_, err = q.CreateAnimeMovieResource(ctx, rarg)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		link, err := q.CreateAnimeLink(ctx, arg.CreateAnimeLinkParams)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		larg := CreateAnimeMovieLinkParams{
			AnimeID: anime.ID,
			LinkID:  link.ID,
		}

		_, err = q.CreateAnimeMovieLink(ctx, larg)
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

				arg := CreateAnimeMovieMetaParams{
					AnimeID:    anime.ID,
					LanguageID: lang.ID,
					MetaID:     meta.ID,
				}
				animeMeta, err := q.CreateAnimeMovieMeta(ctx, arg)
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

		result.AnimeMovie = anime
		result.AnimeResource = resource
		result.AnimeLink = link

		return err
	})

	return result, err
}
