-- name: CreateAnimeMovie :one
INSERT INTO anime_movies (
    original_title,
    aired,
    release_year,
    duration
)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAnimeMovie :one
SELECT * FROM anime_movies 
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeMovie :one
UPDATE anime_movies
SET
  original_title = COALESCE(sqlc.narg(original_title), original_title),
  aired = COALESCE(sqlc.narg(aired), aired),
  release_year = COALESCE(sqlc.narg(release_year), release_year),
  duration = COALESCE(sqlc.narg(duration), duration)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeMovie :exec
DELETE FROM anime_movies 
WHERE id = $1;

-- name: ListAnimeMovies :many
SELECT * FROM anime_movies
WHERE release_year = $1 OR $1 = 0
LIMIT $2
OFFSET $3;