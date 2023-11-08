-- name: CreateAnimeSerieResource :one
INSERT INTO anime_serie_resources (anime_id, resource_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieResource :one
SELECT * FROM anime_serie_resources
WHERE id = $1 LIMIT 1;

-- name: GetAnimeSerieResourceByAnimeID :one
SELECT * FROM anime_serie_resources
WHERE anime_id = $1 LIMIT 1;

-- name: ListAnimeSerieResources :many
SELECT resource_id
FROM anime_serie_resources
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeSerieResource :exec
DELETE FROM anime_serie_resources
WHERE anime_id = $1 AND resource_id = $2;