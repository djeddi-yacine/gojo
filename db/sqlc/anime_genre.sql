-- name: CreateAnimeGenre :exec
INSERT INTO anime_genres (anime_id, genre_id)
VALUES ($1, $2);

-- name: UpdateAnimeGenre :one
UPDATE anime_genres
SET genre_id = $2
WHERE anime_id = $1
RETURNING * ;

-- name: ListAnimeGenres :many
SELECT genre_id
FROM anime_genres
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeGenre :exec
DELETE FROM anime_genres
WHERE anime_id = $1 AND genre_id = $2;