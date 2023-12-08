-- name: CreateAnimeSeasonResource :one
INSERT INTO anime_season_resources (season_id, resource_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSeasonResource :one
SELECT * FROM anime_season_resources
WHERE id = $1;

-- name: ListAnimeSeasonResourcesByAnimeID :many
SELECT resource_id
FROM anime_season_resources
WHERE season_id = $1
ORDER BY id;

-- name: DeleteAnimeSeasonResource :exec
DELETE FROM anime_season_resources
WHERE season_id = $1 AND resource_id = $2;