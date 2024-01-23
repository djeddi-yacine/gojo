-- name: CreateAnimeImage :one
INSERT INTO anime_images (image_host, image_url, image_thumbnails, image_blur_hash, image_height, image_width)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING  *;

-- name: GetAnimeImage :one
SELECT * FROM anime_images
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeImage :one
UPDATE anime_images
SET
  image_host = COALESCE(sqlc.narg(image_host), image_host),
  image_url = COALESCE(sqlc.narg(image_url), image_url),
  image_thumbnails = COALESCE(sqlc.narg(image_thumbnails), image_thumbnails),
  image_blur_hash = COALESCE(sqlc.narg(image_blur_hash), image_blur_hash),
  image_height = COALESCE(sqlc.narg(image_height), image_height),
  image_width = COALESCE(sqlc.narg(image_width), image_width)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: ListAnimeImages :many
SELECT * FROM anime_images
WHERE id = $1;

-- name: DeleteAnimeImage :exec
DELETE FROM anime_images
WHERE id = $1;