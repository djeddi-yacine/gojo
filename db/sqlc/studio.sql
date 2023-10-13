-- name: CreateStudio :one
INSERT INTO studios (studio_name)
VALUES ($1)
RETURNING  id, studio_name, created_at;

-- name: GetStudio :one
SELECT * FROM studios
WHERE id = $1 LIMIT 1;

-- name: UpdateStudio :one
UPDATE studios
SET studio_name = $2
WHERE id = $1
RETURNING id, studio_name, created_at;

-- name: ListStudios :many
SELECT * FROM studios
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteStudio :exec
DELETE FROM studios
WHERE id = $1;
