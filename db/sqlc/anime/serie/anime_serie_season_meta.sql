-- name: CreateAnimeSeasonMeta :one
INSERT INTO anime_season_metas (season_id, language_id, meta_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAnimeSeasonMeta :one
SELECT meta_id FROM anime_season_metas
WHERE season_id = $1 AND language_id = $2;

-- name: ListAnimeSeasonMetasBySeason :many
SELECT * FROM anime_season_metas
WHERE season_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAnimeSeasonMeta :one
UPDATE anime_season_metas
SET
  meta_id = COALESCE(sqlc.narg(meta_id), meta_id),
  season_id = COALESCE(sqlc.narg(season_id), season_id),
  language_id = COALESCE(sqlc.narg(language_id), language_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSeasonMeta :exec
DELETE FROM anime_season_metas
WHERE season_id = $1 AND language_id = $2;