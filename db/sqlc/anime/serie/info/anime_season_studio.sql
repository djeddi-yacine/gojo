-- name: CreateAnimeSeasonStudio :one
INSERT INTO anime_season_studios (season_id, studio_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSeasonStudio :one
SELECT * FROM anime_season_studios
WHERE season_id = $1 AND studio_id = $2;

-- name: ListAnimeSeasonStudios :many
SELECT studio_id
FROM anime_season_studios
WHERE season_id = $1
ORDER BY id;

-- name: DeleteAnimeSeasonStudio :exec
DELETE FROM anime_season_studios
WHERE season_id = $1 AND studio_id = $2;