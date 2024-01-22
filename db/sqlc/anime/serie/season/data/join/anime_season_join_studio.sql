-- name: GetAnimeSeasonStudiosDirectly :many
SELECT studios.*
FROM studios
JOIN anime_season_studios ON studios.id = anime_season_studios.studio_id
WHERE anime_season_studios.season_id = $1;
