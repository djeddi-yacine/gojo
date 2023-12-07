-- name: CreateAnimeSerieImage :one
INSERT INTO anime_serie_images (anime_id, image_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieImage :one
SELECT * FROM anime_serie_images
WHERE id = $1 LIMIT 1;

-- name: GetAnimeSerieImageByAnimeID :one
SELECT * FROM anime_serie_images
WHERE anime_id = $1 LIMIT 1;

-- name: ListAnimeSerieImages :many
SELECT image_id
FROM anime_serie_images
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeSerieImage :exec
DELETE FROM anime_serie_images
WHERE anime_id = $1 AND image_id = $2;