-- name: CreateGenre :one
INSERT INTO genres (genre_name)
VALUES ($1)
RETURNING  *;

-- name: GetGenre :one
SELECT * FROM genres
WHERE id = $1 LIMIT 1;

-- name: UpdateGenre :one
UPDATE genres
SET genre_name = $2
WHERE id = $1
RETURNING *;

-- name: ListGenres :many
SELECT * FROM genres
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteGenre :exec
DELETE FROM genres
WHERE id = $1;
