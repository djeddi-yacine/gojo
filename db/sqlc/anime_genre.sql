-- name: CreateAnimeGenre :one
INSERT INTO anime_genre (anime_id, genre_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeGenre :one
SELECT * FROM anime_genre
WHERE id = $1 LIMIT 1;

-- name: ListAnimeGenres :many
SELECT genre_id
FROM anime_genre
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeGenre :exec
DELETE FROM anime_genre
WHERE anime_id = $1 AND genre_id = $2;