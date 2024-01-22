-- name: CreateAnimeSerieSeason :one
INSERT INTO anime_serie_seasons (anime_id, season_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieSeason :one
SELECT * FROM anime_serie_seasons
WHERE id = $1
LIMIT 1;

-- name: ListAnimeSerieSeasons :many
SELECT season_id FROM anime_serie_seasons
WHERE anime_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAnimeSerieSeason :one
UPDATE anime_serie_seasons
SET
  anime_id = COALESCE(sqlc.narg(anime_id), anime_id),
  season_id = COALESCE(sqlc.narg(season_id), season_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSerieSeason :exec
DELETE FROM anime_serie_seasons
WHERE id = $1;