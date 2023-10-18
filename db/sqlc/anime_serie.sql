-- name: CreateAnimeSerie :one
INSERT INTO anime_series (
    original_title,
    aired,
    release_year,
    duration
)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: GetAnimeSerie :one
SELECT * FROM anime_series 
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeSerie :one
UPDATE anime_series
SET
  original_title = COALESCE(sqlc.narg(original_title), original_title),
  aired = COALESCE(sqlc.narg(aired), aired),
  release_year = COALESCE(sqlc.narg(release_year), release_year),
  duration = COALESCE(sqlc.narg(duration), duration)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSerie :exec
DELETE FROM anime_series 
WHERE id = $1;

-- name: ListAnimeSeries :many
SELECT * FROM anime_series
WHERE release_year = $1 OR $1 = 0
LIMIT $2
OFFSET $3;