-- name: GetAnimeSeasonCharactersDirectly :many
SELECT anime_characters.*
FROM anime_characters
JOIN anime_season_characters ON anime_characters.id = anime_season_characters.studio_id
WHERE anime_season_characters.season_id = $1;
