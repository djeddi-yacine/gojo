-- name: CreateAnimeSerieGenre :one
INSERT INTO anime_serie_genre (anime_id, genre_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieGenre :one
SELECT * FROM anime_serie_genre
WHERE id = $1 LIMIT 1;

-- name: ListAnimeSerieGenres :many
SELECT genre_id
FROM anime_serie_genre
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeSerieGenre :exec
DELETE FROM anime_serie_genre
WHERE anime_id = $1 AND genre_id = $2;