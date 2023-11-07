-- name: CreateAnimeSerieSeasonMeta :one
INSERT INTO anime_serie_season_metas (season_id, language_id, meta_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAnimeSerieSeasonMeta :one
SELECT * FROM anime_serie_season_metas
WHERE season_id = $1 AND language_id = $2;

-- name: ListAnimeSerieSeasonMetasBySeason :many
SELECT * FROM anime_serie_season_metas
WHERE season_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAnimeSerieSeasonMeta :one
UPDATE anime_serie_season_metas
SET
  meta_id = COALESCE(sqlc.narg(meta_id), meta_id),
  season_id = COALESCE(sqlc.narg(season_id), season_id),
  language_id = COALESCE(sqlc.narg(language_id), language_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSerieSeasonMeta :exec
DELETE FROM anime_serie_season_metas
WHERE season_id = $1 AND language_id = $2;