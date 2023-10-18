-- name: CreateAnimeMovie :one
INSERT INTO anime_movie (
    original_title,
    aired,
    release_year,
    duration
)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAnimeMovie :one
SELECT * FROM anime_movie 
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeMovie :one
UPDATE anime_movie
SET
  original_title = COALESCE($2, original_title),
  aired = COALESCE($3, aired),
  release_year = COALESCE($4, release_year),
  duration = COALESCE($5, duration)
WHERE
  id = $1
RETURNING *;

-- name: DeleteAnimeMovie :exec
DELETE FROM anime_movie 
WHERE id = $1;

-- name: ListAnimeMovies :many
SELECT * FROM anime_movie
WHERE release_year = $1 OR $1 = 0
LIMIT $2
OFFSET $3;