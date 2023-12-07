-- name: CreateAnimeMovieImage :one
INSERT INTO anime_movie_images (anime_id, image_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieImage :one
SELECT * FROM anime_movie_images
WHERE id = $1 LIMIT 1;

-- name: GetAnimeMovieImageByAnimeID :one
SELECT * FROM anime_movie_images
WHERE anime_id = $1 LIMIT 1;

-- name: ListAnimeMovieImages :many
SELECT image_id
FROM anime_movie_images
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeMovieImage :exec
DELETE FROM anime_movie_images
WHERE anime_id = $1 AND image_id = $2;