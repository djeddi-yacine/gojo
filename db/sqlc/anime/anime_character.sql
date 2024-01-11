-- name: CreateAnimeCharacter :one
INSERT INTO anime_characters (full_name, about, role_playing, image_url, image_blur_hash, pictures)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (full_name, about)
DO UPDATE SET 
    pictures = ARRAY(SELECT DISTINCT UNNEST(anime_characters.pictures || excluded.pictures))
RETURNING *;

-- name: GetAnimeCharacter :one
SELECT * FROM anime_characters
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeCharacter :one
UPDATE anime_characters
SET
  full_name = COALESCE(sqlc.narg(full_name), full_name),
  about = COALESCE(sqlc.narg(about), about),
  role_playing = COALESCE(sqlc.narg(role_playing), role_playing),
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