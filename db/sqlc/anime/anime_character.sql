-- name: CreateAnimeCharacter :one
INSERT INTO anime_characters (actor_id, full_name, about, image_url, image_blur_hash)
VALUES ($1, $2, $3, $4, $5)
RETURNING  *;

-- name: GetAnimeCharacter :one
SELECT * FROM anime_characters
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeCharacter :one
UPDATE anime_characters
SET
  actor_id = COALESCE(sqlc.narg(actor_id), actor_id),
  full_name = COALESCE(sqlc.narg(full_name), full_name),
  about = COALESCE(sqlc.narg(about), about),
  image_url = COALESCE(sqlc.narg(image_url), image_url),
  image_blur_hash = COALESCE(sqlc.narg(image_blur_hash), image_blur_hash)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: ListAnimeCharacters :many
SELECT * FROM anime_characters
WHERE id = $1;

-- name: DeleteAnimeCharacter :exec
DELETE FROM anime_characters
WHERE id = $1;