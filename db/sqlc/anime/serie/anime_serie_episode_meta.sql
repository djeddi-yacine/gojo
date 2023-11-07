-- name: CreateAnimeSerieEpisodeMeta :one
INSERT INTO anime_serie_episode_metas (episode_id, meta_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieEpisodeMeta :one
SELECT * FROM anime_serie_episode_metas
WHERE id = $1
LIMIT 1;

-- name: ListAnimeSerieEpisodeMetasByEpisode :many
SELECT * FROM anime_serie_episode_metas
WHERE episode_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAnimeSerieEpisodeMeta :one
UPDATE anime_serie_episode_metas
SET
  meta_id = COALESCE(sqlc.narg(meta_id), meta_id),
  episode_id = COALESCE(sqlc.narg(episode_id), episode_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSerieEpisodeMeta :exec
DELETE FROM anime_serie_episode_metas
WHERE id = $1;