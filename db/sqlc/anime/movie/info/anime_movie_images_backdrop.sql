-- name: CreateAnimeMovieBackdropImage :one
INSERT INTO anime_movie_backdrop_images (anime_id, image_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListAnimeMovieBackdropImages :many
SELECT image_id
FROM anime_movie_backdrop_images
WHERE anime_id = $1;

-- name: DeleteAnimeMovieBackdropImage :exec
DELETE FROM anime_movie_backdrop_images
WHERE anime_id = $1 AND image_id = $2;