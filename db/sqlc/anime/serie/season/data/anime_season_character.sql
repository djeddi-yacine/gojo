-- name: CreateAnimeSeasonCharacter :one
INSERT INTO anime_season_characters (season_id, character_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSeasonCharacter :one
SELECT * FROM anime_season_characters
WHERE season_id = $1 AND character_id = $2;

-- name: ListAnimeSeasonCharacters :many
SELECT character_id
FROM anime_season_characters
WHERE season_id = $1
ORDER BY id;

-- name: DeleteAnimeSeasonCharacter :exec
DELETE FROM anime_season_characters
WHERE season_id = $1 AND character_id = $2;