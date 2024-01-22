-- name: CreateAnimeSerieBackdropImage :one
INSERT INTO anime_serie_backdrop_images (anime_id, image_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListAnimeSerieBackdropImages :many
SELECT image_id
FROM anime_serie_backdrop_images
WHERE anime_id = $1;

-- name: DeleteAnimeSerieBackdropImage :exec
DELETE FROM anime_serie_backdrop_images
WHERE anime_id = $1 AND image_id = $2;