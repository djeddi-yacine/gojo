-- name: CreateAnimeSerieResource :one
INSERT INTO anime_serie_resources (anime_id, resource_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieResource :one
SELECT * FROM anime_serie_resources
WHERE id = $1;

-- name: ListAnimeSerieResourcesByAnimeID :many
SELECT resource_id
FROM anime_serie_resources
WHERE anime_id = $1
ORDER BY id;

-- name: DeleteAnimeSerieResource :exec
DELETE FROM anime_serie_resources
WHERE anime_id = $1 AND resource_id = $2;