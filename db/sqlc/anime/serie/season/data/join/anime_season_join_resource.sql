-- name: GetAnimeSeasonResourceDirectly :one
SELECT anime_resources.*
FROM anime_resources
JOIN anime_season_resources ON anime_resources.id = anime_season_resources.resource_id
WHERE anime_season_resources.season_id = $1
LIMIT 1;