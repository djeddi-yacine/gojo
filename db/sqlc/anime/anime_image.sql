-- name: CreateAnimeImage :one
INSERT INTO anime_images (image_type, image_host, image_url, image_thumbnails, image_blurhash, image_height, image_width)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING  *;

-- name: GetAnimeImage :one
SELECT * FROM anime_images
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeImage :one
UPDATE anime_images
SET
  image_type = COALESCE(sqlc.narg(image_type), image_type),
  image_host = COALESCE(sqlc.narg(image_host), image_host),
  image_url = COALESCE(sqlc.narg(image_url), image_url),
  image_thumbnails = COALESCE(sqlc.narg(image_thumbnails), image_thumbnails),
  image_blurhash = COALESCE(sqlc.narg(image_blurhash), image_blurhash),
  image_height = COALESCE(sqlc.narg(image_height), image_height),
  image_width = COALESCE(sqlc.narg(image_width), image_width)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeImage :exec
DELETE FROM anime_images
WHERE id = $1;