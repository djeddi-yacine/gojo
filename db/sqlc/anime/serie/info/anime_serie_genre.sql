-- name: CreateAnimeSerieGenre :one
INSERT INTO anime_serie_genres (anime_id, genre_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieGenre :one
SELECT * FROM anime_serie_genres
WHERE anime_id = $1 AND genre_id = $2;

-- name: ListAnimeSerieGenres :many
SELECT genre_id
FROM anime_serie_genres
WHERE anime_id = $1
ORDER BY id;

-- name: DeleteAnimeSerieGenre :exec
DELETE FROM anime_serie_genres
WHERE anime_id = $1 AND genre_id = $2;