-- name: CreateAnimeSerieSeason :one
INSERT INTO anime_serie_seasons (anime_id, season_number)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieSeason :one
SELECT * FROM anime_serie_seasons
WHERE id = $1
LIMIT 1;

-- name: ListAnimeSerieSeasonsByAnime :many
SELECT * FROM anime_serie_seasons
WHERE anime_id = $1
ORDER BY season_number
LIMIT $2
OFFSET $3;

-- name: UpdateAnimeSerieSeason :one
UPDATE anime_serie_seasons
SET season_number = $2
WHERE id = $1
RETURNING *;

-- name: DeleteAnimeSerieSeason :exec
DELETE FROM anime_serie_seasons
WHERE id = $1;
