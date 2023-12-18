-- name: CreateAnimeSeasonTag :one
INSERT INTO anime_season_tags (season_id, tag_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListAnimeSeasonTags :many
SELECT * FROM anime_season_tags
WHERE season_id = $1;

-- name: DeleteAnimeSeasonTag :exec
DELETE FROM anime_season_tags
WHERE season_id = $1 AND tag_id = $2;