-- name: CreateAnimeSerieCharacter :one
INSERT INTO anime_serie_characters (anime_id, character_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieCharacter :one
SELECT * FROM anime_serie_characters
WHERE anime_id = $1 AND character_id = $2;

-- name: ListAnimeSerieCharacters :many
SELECT character_id
FROM anime_serie_characters
WHERE anime_id = $1
ORDER BY id;

-- name: DeleteAnimeSerieCharacter :exec
DELETE FROM anime_serie_characters
WHERE anime_id = $1 AND character_id = $2;