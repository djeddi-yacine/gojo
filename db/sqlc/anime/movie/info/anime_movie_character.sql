-- name: CreateAnimeMovieCharacter :one
INSERT INTO anime_movie_characters (anime_id, character_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieCharacter :one
SELECT * FROM anime_movie_characters
WHERE anime_id = $1 AND character_id = $2;

-- name: ListAnimeMovieCharacters :many
SELECT character_id
FROM anime_movie_characters
WHERE anime_id = $1
ORDER BY id;

-- name: DeleteAnimeMovieCharacter :exec
DELETE FROM anime_movie_characters
WHERE anime_id = $1 AND character_id = $2;