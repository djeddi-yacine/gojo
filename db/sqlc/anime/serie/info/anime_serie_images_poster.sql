-- name: CreateAnimeSeriePosterImage :one
INSERT INTO anime_serie_poster_images (anime_id, image_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSeriePosterImage :one
SELECT * FROM anime_serie_poster_images
WHERE id = $1 LIMIT 1;

-- name: GetAnimeSeriePosterImageByAnimeID :one
SELECT * FROM anime_serie_poster_images
WHERE anime_id = $1 LIMIT 1;

-- name: ListAnimeSeriePosterImages :many
SELECT image_id
FROM anime_serie_poster_images
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeSeriePosterImage :exec
DELETE FROM anime_serie_poster_images
WHERE anime_id = $1 AND image_id = $2;