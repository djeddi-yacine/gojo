package db

import (
	"context"
)

func (gojo *SQLGojo) CreateAnimeEpisodeServerTx(ctx context.Context, episodeID int64) (AnimeEpisodeServer, error) {
	var result AnimeEpisodeServer

	err := gojo.execTx(ctx, func(q *Queries) error {
		var err error

		_, err = q.GetAnimeEpisodeByEpisodeID(ctx, episodeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		server, err := q.CreateAnimeEpisodeServer(ctx, episodeID)
		if err != nil {
			ErrorSQL(err)
			return err
		}

		result = server

		return err
	})

	return result, err
}
