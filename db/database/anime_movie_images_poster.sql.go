// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.23.0
// source: anime_movie_images_poster.sql

package db

import (
	"context"
)

const createAnimeMoviePosterImage = `-- name: CreateAnimeMoviePosterImage :one
INSERT INTO anime_movie_poster_images (anime_id, image_id)
VALUES ($1, $2)
RETURNING id, anime_id, image_id, created_at
`

type CreateAnimeMoviePosterImageParams struct {
	AnimeID int64
	ImageID int64
}

func (q *Queries) CreateAnimeMoviePosterImage(ctx context.Context, arg CreateAnimeMoviePosterImageParams) (AnimeMoviePosterImage, error) {
	row := q.db.QueryRow(ctx, createAnimeMoviePosterImage, arg.AnimeID, arg.ImageID)
	var i AnimeMoviePosterImage
	err := row.Scan(
		&i.ID,
		&i.AnimeID,
		&i.ImageID,
		&i.CreatedAt,
	)
	return i, err
}

const deleteAnimeMoviePosterImage = `-- name: DeleteAnimeMoviePosterImage :exec
DELETE FROM anime_movie_poster_images
WHERE anime_id = $1 AND image_id = $2
`

type DeleteAnimeMoviePosterImageParams struct {
	AnimeID int64
	ImageID int64
}

func (q *Queries) DeleteAnimeMoviePosterImage(ctx context.Context, arg DeleteAnimeMoviePosterImageParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeMoviePosterImage, arg.AnimeID, arg.ImageID)
	return err
}

const listAnimeMoviePosterImages = `-- name: ListAnimeMoviePosterImages :many
SELECT image_id
FROM anime_movie_poster_images
WHERE anime_id = $1
`

func (q *Queries) ListAnimeMoviePosterImages(ctx context.Context, animeID int64) ([]int64, error) {
	rows, err := q.db.Query(ctx, listAnimeMoviePosterImages, animeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int64{}
	for rows.Next() {
		var image_id int64
		if err := rows.Scan(&image_id); err != nil {
			return nil, err
		}
		items = append(items, image_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}