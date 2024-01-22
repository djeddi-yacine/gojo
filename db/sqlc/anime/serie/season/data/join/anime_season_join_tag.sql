-- name: GetAnimeSeasonTagsDirectly :many
SELECT anime_tags.*
FROM anime_tags
JOIN anime_season_tags ON anime_tags.id = anime_season_tags.tag_id
WHERE anime_season_tags.season_id = $1;
