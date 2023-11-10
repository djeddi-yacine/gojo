package db

import (
	"context"
	"errors"
)

type AddAnimeSerieEpisodeMetasTxParams struct {
	EpisodeID    int64
	EpisodeMetas []AnimeMetaTxParam
}

type AddAnimeSerieEpisodeMetasTxResult struct {
	AnimeSerieEpisode      AnimeSerieEpisode
	AnimeSerieEpisodeMetas []AnimeMetaTxResult
}

func (gojo *SQLGojo) AddAnimeSerieEpisodeMetasTx(ctx context.Context, arg AddAnimeSerieEpisodeMetasTxParams) (AddAnimeSerieEpisodeMetasTxResult, error) {
	var result AddAnimeSerieEpisodeMetasTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		Episode, err := q.GetAnimeSerieEpisode(ctx, arg.EpisodeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeSerieEpisode = Episode

		if arg.EpisodeMetas != nil {
			var metaArg CreateMetaParams
			var EpisodeMetaArg CreateAnimeSerieEpisodeMetaParams
			result.AnimeSerieEpisodeMetas = make([]AnimeMetaTxResult, len(arg.EpisodeMetas))

			for i, m := range arg.EpisodeMetas {
				metaArg = CreateMetaParams{
					Title:    m.Title,
					Overview: m.Overview,
				}

				meta, err := q.CreateMeta(ctx, metaArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				EpisodeMetaArg = CreateAnimeSerieEpisodeMetaParams{
					EpisodeID:  Episode.ID,
					LanguageID: m.LanguageID,
					MetaID:     meta.ID,
				}

				_, err = q.CreateAnimeSerieEpisodeMeta(ctx, EpisodeMetaArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}
				
				result.AnimeSerieEpisodeMetas[i] = AnimeMetaTxResult{
					Meta:     meta,
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
