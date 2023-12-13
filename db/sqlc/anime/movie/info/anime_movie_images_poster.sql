-- name: CreateAnimeMoviePosterImage :one
INSERT INTO anime_movie_poster_images (anime_id, image_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMoviePosterImage :one
SELECT * FROM anime_movie_poster_images
WHERE id = $1 LIMIT 1;

-- name: GetAnimeMoviePosterImageByAnimeID :one
SELECT * FROM anime_movie_poster_images
WHERE anime_id = $1 LIMIT 1;

-- name: ListAnimeMoviePosterImages :many
SELECT image_id
FROM anime_movie_poster_images
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeMoviePosterImage :exec
DELETE FROM anime_movie_poster_images
WHERE anime_id = $1 AND image_id = $2;