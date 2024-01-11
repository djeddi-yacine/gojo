-- name: CreateLanguage :one
INSERT INTO languages (language_name, language_code)
VALUES ($1, $2)
ON CONFLICT (language_code)
DO UPDATE SET language_code = excluded.language_code
RETURNING  *;

-- name: GetLanguage :one
SELECT * FROM languages
WHERE id = $1 LIMIT 1;

-- name: UpdateLanguage :one
UPDATE languages
SET language_code = $2,
    language_name = $3
WHERE id = $1
RETURNING *;

-- name: ListLanguages :many
SELECT id FROM languages
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteLanguage :exec
DELETE FROM languages
WHERE id = $1;