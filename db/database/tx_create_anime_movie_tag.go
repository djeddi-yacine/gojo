package db

import (
	"context"

	"github.com/jackc/pgerrcode"
)

type CreateAnimeMovieTagTxParams struct {
	AnimeID   int64
	AnimeTags []string
}

type CreateAnimeMovieTagTxResult struct {
	AnimeMovie AnimeMovie
	AnimeTags  []AnimeTag
}

func (gojo *SQLGojo) CreateAnimeMovieTagTx(ctx context.Context, arg CreateAnimeMovieTagTxParams) (CreateAnimeMovieTagTxResult, error) {
	var result CreateAnimeMovieTagTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		anime, err := q.GetAnimeMovie(ctx, arg.AnimeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeMovie = anime

		if arg.AnimeTags != nil {
			if len(arg.AnimeTags) > 0 {
				var tag AnimeTag
				tagsArg := make([]CreateAnimeMovieTagParams, len(arg.AnimeTags))
				result.AnimeTags = make([]AnimeTag, len(arg.AnimeTags))

				for i, t := range arg.AnimeTags {
					tag, err = q.CreateAnimeTag(ctx, t)
					if err != nil {
						if ErrorDB(err).Code == pgerrcode.UniqueViolation {
							tag, err = q.GetAnimeTagByTag(ctx, t)
							if err != nil {
								ErrorSQL(err)
								return err
							}
						} else {
							ErrorSQL(err)
							return err
						}
					}

					result.AnimeTags[i] = tag
					tagsArg[i].AnimeID = anime.ID
					tagsArg[i].TagID = tag.ID
				}

				for _, amt := range tagsArg {
					_, err = q.CreateAnimeMovieTag(ctx, amt)
					if err != nil {
						if ErrorDB(err).Code != pgerrcode.UniqueViolation {
							ErrorSQL(err)
							return err
						}
					}
				}
			}
		}

		return err
	})

	return result, err
}
