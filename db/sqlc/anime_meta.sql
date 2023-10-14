-- name: CreateAnimeMeta :one
INSERT INTO anime_metas (anime_id, language_id, meta_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAnimeMeta :one
SELECT meta_id
FROM anime_metas
WHERE anime_id = $1 AND language_id = $2;

-- name: UpdateAnimeMeta :one
UPDATE anime_metas
SET meta_id = $3
WHERE anime_id = $1 AND language_id = $2
RETURNING * ;

-- name: ListAnimeMetas :many
SELECT meta_id
FROM anime_metas
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeMeta :exec
DELETE FROM anime_metas
WHERE anime_id = $1 AND language_id = $2;