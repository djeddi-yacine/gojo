-- name: CreateAnimeSerieLogoImage :one
INSERT INTO anime_serie_logo_images (anime_id, image_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListAnimeSerieLogoImages :many
SELECT image_id
FROM anime_serie_logo_images
WHERE anime_id = $1;

-- name: DeleteAnimeSerieLogoImage :exec
DELETE FROM anime_serie_logo_images
WHERE anime_id = $1 AND image_id = $2;