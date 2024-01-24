-- name: CreateStudio :one
INSERT INTO studios (studio_name)
VALUES ($1)
ON CONFLICT (studio_name)
DO UPDATE SET studio_name = excluded.studio_name
RETURNING  *;

-- name: GetStudio :one
SELECT * FROM studios
WHERE id = $1 LIMIT 1;

-- name: UpdateStudio :one
UPDATE studios
SET
    studio_name = COALESCE(sqlc.narg(studio_name), studio_name)
WHERE id = $1
RETURNING *;

-- name: ListStudios :many
SELECT id FROM studios
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteStudio :exec
DELETE FROM studios
WHERE id = $1;
