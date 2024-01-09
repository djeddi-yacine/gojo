-- name: CreateAnimeMovie :one
INSERT INTO anime_movies (
    original_title,
    unique_id,
    aired,
    release_year,
    rating,
    duration,
    portrait_poster,
    portrait_blur_hash,
    landscape_poster,
    landscape_blur_hash
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
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
  rating = COALESCE(sqlc.narg(rating), rating),
  duration = COALESCE(sqlc.narg(duration), duration),
  portrait_poster = COALESCE(sqlc.narg(portrait_poster), portrait_poster),
  portrait_blur_hash = COALESCE(sqlc.narg(portrait_blur_hash), portrait_blur_hash),
  landscape_poster = COALESCE(sqlc.narg(landscape_poster), landscape_poster),
  landscape_blur_hash = COALESCE(sqlc.narg(landscape_blur_hash), landscape_blur_hash)
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