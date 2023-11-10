-- name: CreateAnimeSerieMeta :one
INSERT INTO anime_serie_metas (anime_id, language_id, meta_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAnimeSerieMeta :one
SELECT meta_id
FROM anime_serie_metas
WHERE anime_id = $1 AND language_id = $2;

-- name: UpdateAnimeSerieMeta :one
UPDATE anime_serie_metas
SET meta_id = $3
WHERE anime_id = $1 AND language_id = $2
RETURNING * ;

-- name: ListAnimeSerieMetas :many
SELECT meta_id
FROM anime_serie_metas
WHERE anime_id = $1
ORDER BY id;

-- name: DeleteAnimeSerieMeta :exec
DELETE FROM anime_serie_metas
WHERE anime_id = $1 AND language_id = $2;