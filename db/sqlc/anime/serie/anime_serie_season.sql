-- name: CreateAnimeSerieSeason :one
INSERT INTO anime_serie_seasons (
    anime_id,
    season_number,
    portriat_poster,
    portriat_blur_hash,
    landscape_poster,
    landscape_blur_hash
)
VALUES ($1, $2, $3, $4, $5, $6)
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
SET
  season_number = COALESCE(sqlc.narg(season_number), season_number),
  portriat_poster = COALESCE(sqlc.narg(portriat_poster), portriat_poster),
  portriat_blur_hash = COALESCE(sqlc.narg(portriat_blur_hash), portriat_blur_hash),
  landscape_poster = COALESCE(sqlc.narg(landscape_poster), landscape_poster),
  landscape_blur_hash = COALESCE(sqlc.narg(landscape_blur_hash), landscape_blur_hash)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSerieSeason :exec
DELETE FROM anime_serie_seasons
WHERE id = $1;
