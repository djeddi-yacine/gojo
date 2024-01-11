-- name: CreateActor :one
INSERT INTO actors (full_name, gender, biography, born, image_url, image_blur_hash)
VALUES ($1, $2, $3, $4, $5, $6)
ON CONFLICT (full_name, born)
DO UPDATE SET gender = excluded.gender
RETURNING  *;

-- name: GetActor :one
SELECT * FROM actors
WHERE id = $1 LIMIT 1;

-- name: UpdateActor :one
UPDATE actors
SET 
  full_name = COALESCE(sqlc.narg(full_name), full_name),
  gender = COALESCE(sqlc.narg(gender), gender),
  biography = COALESCE(sqlc.narg(biography), biography),
  born = COALESCE(sqlc.narg(born), born),
  image_url = COALESCE(sqlc.narg(image_url), image_url),
  image_blur_hash = COALESCE(sqlc.narg(image_blur_hash), image_blur_hash)
WHERE id = $1
RETURNING *;

-- name: ListActors :many
SELECT id FROM actors
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteActor :exec
DELETE FROM actors
WHERE id = $1;
