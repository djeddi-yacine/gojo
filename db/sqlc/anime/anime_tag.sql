-- name: CreateAnimeTag :one
INSERT INTO anime_tags (tag) 
VALUES ($1)
RETURNING  *;

-- name: GetAnimeTag :one
SELECT * FROM anime_tags
WHERE id = $1 LIMIT 1;

-- name: GetAnimeTagByTag :one
SELECT * FROM anime_tags
WHERE tag = $1 LIMIT 1;

-- name: UpdateAnimeTag :one
UPDATE anime_tags
SET
  tag = COALESCE(sqlc.narg(tag), tag)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeTag :exec
DELETE FROM anime_tags
WHERE id = $1;