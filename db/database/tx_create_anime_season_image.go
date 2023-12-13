package db

import (
	"context"
)

type CreateAnimeSeasonImageTxParams struct {
	SeasonID      int64
	SeasonPosters []CreateAnimeImageParams
}

type CreateAnimeSeasonImageTxResult struct {
	AnimeSeason  AnimeSerieSeason
	AnimePosters []AnimeImage
}

func (gojo *SQLGojo) CreateAnimeSeasonImageTx(ctx context.Context, arg CreateAnimeSeasonImageTxParams) (CreateAnimeSeasonImageTxResult, error) {
	var result CreateAnimeSeasonImageTxResult

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		season, err := q.GetAnimeSerieSeason(ctx, arg.SeasonID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result.AnimeSeason = season

		if arg.SeasonPosters != nil {
			if len(arg.SeasonPosters) > 0 {
				var imgArg CreateAnimeImageParams
				postersArg := make([]CreateAnimeSeasonPosterImageParams, len(arg.SeasonPosters))
				result.AnimePosters = make([]AnimeImage, len(arg.SeasonPosters))

				for i, p := range arg.SeasonPosters {
					imgArg = CreateAnimeImageParams{
						ImageHost:       p.ImageHost,
						ImageUrl:        p.ImageUrl,
						ImageThumbnails: p.ImageThumbnails,
						ImageBlurhash:   p.ImageBlurhash,
						ImageHeight:     p.ImageHeight,
						ImageWidth:      p.ImageWidth,
					}

					img, err := q.CreateAnimeImage(ctx, imgArg)
					if err != nil {
						ErrorSQL(err)
						return err
					}

					result.AnimePosters[i] = img
					postersArg[i].SeasonID = season.ID
					postersArg[i].ImageID = img.ID
				}

				for _, asp := range postersArg {
					_, err = q.CreateAnimeSeasonPosterImage(ctx, asp)
					if err != nil {
						ErrorSQL(err)
						return err
					}
				}
			}
		}

		return err
	})

	return result, err
}
