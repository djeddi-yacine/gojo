-- name: CreateAnimeSeasonGenre :one
INSERT INTO anime_season_genres (season_id, genre_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSeasonGenre :one
SELECT * FROM anime_season_genres
WHERE season_id = $1 AND genre_id = $2;

-- name: ListAnimeSeasonGenres :many
SELECT genre_id
FROM anime_season_genres
WHERE season_id = $1
ORDER BY id;

-- name: DeleteAnimeSeasonGenre :exec
DELETE FROM anime_season_genres
WHERE season_id = $1 AND genre_id = $2;