-- name: CreateAnimeSeason :one
INSERT INTO anime_serie_seasons (
    anime_id,
    season_original_title,
    release_year,
    aired,
    rating,
    portrait_poster,
    portrait_blur_hash
)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetAnimeSeason :one
SELECT * FROM anime_serie_seasons
WHERE id = $1
LIMIT 1;

-- name: ListAnimeSeasonsByAnimeID :many
SELECT * FROM anime_serie_seasons
WHERE anime_id = $1
ORDER BY release_year
LIMIT $2
OFFSET $3;

-- name: UpdateAnimeSeason :one
UPDATE anime_serie_seasons
SET
  season_original_title = COALESCE(sqlc.narg(season_original_title), season_original_title),
  release_year = COALESCE(sqlc.narg(release_year), release_year),
  aired = COALESCE(sqlc.narg(aired), aired),
  rating = COALESCE(sqlc.narg(rating), rating),
  portrait_poster = COALESCE(sqlc.narg(portrait_poster), portrait_poster),
  portrait_blur_hash = COALESCE(sqlc.narg(portrait_blur_hash), portrait_blur_hash)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSeason :exec
DELETE FROM anime_serie_seasons
WHERE id = $1;
