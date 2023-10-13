-- name: CreateMeta :one
INSERT INTO metas (title, overview)
VALUES ($1, $2)
RETURNING  *;

-- name: GetMeta :one
SELECT * FROM metas
WHERE id = $1 LIMIT 1;

-- name: UpdateMeta :exec
UPDATE metas
SET title = $2,
    overview = $3
WHERE id = $1;

-- name: DeleteMeta :exec
DELETE FROM Metas
WHERE id = $1;
