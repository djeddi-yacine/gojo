-- name: GetAnimeMovie :one
SELECT * FROM anime_movie 
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeMovie :one
UPDATE anime_movie
SET
  original_title = COALESCE(sqlc.narg(original_title), original_title),
  aired = COALESCE(sqlc.narg(aired), aired),
  release_year = COALESCE(sqlc.narg(release_year), release_year),
  duration = COALESCE(sqlc.narg(duration), duration)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: CreateAnimeMovie :one
INSERT INTO anime_movie (
    original_title,
    aired,
    release_year,
    duration
)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteAnimeMovie :exec
DELETE FROM anime_movie 
WHERE id = $1;
