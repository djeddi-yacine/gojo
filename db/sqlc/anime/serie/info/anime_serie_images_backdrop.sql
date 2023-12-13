-- name: CreateAnimeSerieBackdropImage :one
INSERT INTO anime_serie_backdrop_images (anime_id, image_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieBackdropImage :one
SELECT * FROM anime_serie_backdrop_images
WHERE id = $1 LIMIT 1;

-- name: GetAnimeSerieBackdropImageByAnimeID :one
SELECT * FROM anime_serie_backdrop_images
WHERE anime_id = $1 LIMIT 1;

-- name: ListAnimeSerieBackdropImages :many
SELECT image_id
FROM anime_serie_backdrop_images
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeSerieBackdropImage :exec
DELETE FROM anime_serie_backdrop_images
WHERE anime_id = $1 AND image_id = $2;