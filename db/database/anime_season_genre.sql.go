// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: anime_season_genre.sql

package db

import (
	"context"
)

const createAnimeSeasonGenre = `-- name: CreateAnimeSeasonGenre :one
INSERT INTO anime_season_genres (season_id, genre_id)
VALUES ($1, $2)
RETURNING id, season_id, genre_id
`

type CreateAnimeSeasonGenreParams struct {
	SeasonID int64
	GenreID  int32
}

func (q *Queries) CreateAnimeSeasonGenre(ctx context.Context, arg CreateAnimeSeasonGenreParams) (AnimeSeasonGenre, error) {
	row := q.db.QueryRow(ctx, createAnimeSeasonGenre, arg.SeasonID, arg.GenreID)
	var i AnimeSeasonGenre
	err := row.Scan(&i.ID, &i.SeasonID, &i.GenreID)
	return i, err
}

const deleteAnimeSeasonGenre = `-- name: DeleteAnimeSeasonGenre :exec
DELETE FROM anime_season_genres
WHERE season_id = $1 AND genre_id = $2
`

type DeleteAnimeSeasonGenreParams struct {
	SeasonID int64
	GenreID  int32
}

func (q *Queries) DeleteAnimeSeasonGenre(ctx context.Context, arg DeleteAnimeSeasonGenreParams) error {
	_, err := q.db.Exec(ctx, deleteAnimeSeasonGenre, arg.SeasonID, arg.GenreID)
	return err
}

const getAnimeSeasonGenre = `-- name: GetAnimeSeasonGenre :one
SELECT id, season_id, genre_id FROM anime_season_genres
WHERE season_id = $1 AND genre_id = $2
`

type GetAnimeSeasonGenreParams struct {
	SeasonID int64
	GenreID  int32
}

func (q *Queries) GetAnimeSeasonGenre(ctx context.Context, arg GetAnimeSeasonGenreParams) (AnimeSeasonGenre, error) {
	row := q.db.QueryRow(ctx, getAnimeSeasonGenre, arg.SeasonID, arg.GenreID)
	var i AnimeSeasonGenre
	err := row.Scan(&i.ID, &i.SeasonID, &i.GenreID)
	return i, err
}

const listAnimeSeasonGenres = `-- name: ListAnimeSeasonGenres :many
SELECT genre_id
FROM anime_season_genres
WHERE season_id = $1
ORDER BY id
`

func (q *Queries) ListAnimeSeasonGenres(ctx context.Context, seasonID int64) ([]int32, error) {
	rows, err := q.db.Query(ctx, listAnimeSeasonGenres, seasonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []int32{}
	for rows.Next() {
		var genre_id int32
		if err := rows.Scan(&genre_id); err != nil {
			return nil, err
		}
		items = append(items, genre_id)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
