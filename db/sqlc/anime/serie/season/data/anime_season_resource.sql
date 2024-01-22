-- name: CreateAnimeSeasonResource :one
INSERT INTO anime_season_resources (season_id, resource_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSeasonResource :one
SELECT * FROM anime_season_resources
WHERE season_id = $1
LIMIT 1;

-- name: DeleteAnimeSeasonResource :exec
DELETE FROM anime_season_resources
WHERE season_id = $1 AND resource_id = $2;