-- name: CreateAnimeEpisodeMeta :one
INSERT INTO anime_episode_metas (episode_id, language_id, meta_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAnimeEpisodeMeta :one
SELECT meta_id FROM anime_episode_metas
WHERE episode_id = $1 AND language_id = $2;

-- name: ListAnimeEpisodeMetasByEpisode :many
SELECT * FROM anime_episode_metas
WHERE episode_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAnimeEpisodeMeta :one
UPDATE anime_episode_metas
SET
  meta_id = COALESCE(sqlc.narg(meta_id), meta_id),
  episode_id = COALESCE(sqlc.narg(episode_id), episode_id),
  language_id = COALESCE(sqlc.narg(language_id), language_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeEpisodeMeta :exec
DELETE FROM anime_episode_metas
WHERE episode_id = $1 AND language_id = $2;