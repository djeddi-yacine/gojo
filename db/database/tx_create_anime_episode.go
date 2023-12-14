package db

import (
	"context"
	"errors"
)

type CreateAnimeEpisodeTxParams struct {
	Episode      CreateAnimeEpisodeParams
	EpisodeMetas []AnimeMetaTxParam
}

type CreateAnimeEpisodeTxResult struct {
	AnimeSeason       AnimeSerieSeason
	AnimeEpisode      AnimeSerieEpisode
	AnimeEpisodeMetas []AnimeMetaTxResult
}

func (gojo *SQLGojo) CreateAnimeEpisodeTx(ctx context.Context, arg CreateAnimeEpisodeTxParams) (CreateAnimeEpisodeTxResult, error) {
	var result CreateAnimeEpisodeTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		season, err := q.GetAnimeSeason(ctx, arg.Episode.SeasonID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		episode, err := q.CreateAnimeEpisode(ctx, arg.Episode)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		_, err = q.CreateAnimeSeasonEpisode(ctx, CreateAnimeSeasonEpisodeParams{
			SeasonID:  season.ID,
			EpisodeID: episode.ID,
		})
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeEpisode = episode
		result.AnimeSeason = season

		if arg.EpisodeMetas != nil {
			var metaArg CreateMetaParams
			var EpisodeMetaArg CreateAnimeEpisodeMetaParams
			result.AnimeEpisodeMetas = make([]AnimeMetaTxResult, len(arg.EpisodeMetas))

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

				EpisodeMetaArg = CreateAnimeEpisodeMetaParams{
					EpisodeID:  episode.ID,
					LanguageID: m.LanguageID,
					MetaID:     meta.ID,
				}

				_, err = q.CreateAnimeEpisodeMeta(ctx, EpisodeMetaArg)
				if err != nil {
					ErrorSQL(err)
					return err
				}

				result.AnimeEpisodeMetas[i] = AnimeMetaTxResult{
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
