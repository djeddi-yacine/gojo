-- name: CreateAnimeSerieEpisodeServer :one
INSERT INTO anime_serie_episode_servers (episode_id, server_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieEpisodeServer :one
SELECT * FROM anime_serie_episode_servers
WHERE id = $1
LIMIT 1;

-- name: ListAnimeSerieEpisodeServersByEpisode :many
SELECT * FROM anime_serie_episode_servers
WHERE episode_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAnimeSerieEpisodeServer :one
UPDATE anime_serie_episode_servers
SET
  episode_id = COALESCE(sqlc.narg(episode_id), episode_id),
  server_id = COALESCE(sqlc.narg(server_id), server_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;
-- name: DeleteAnimeSerieEpisodeServer :exec
DELETE FROM anime_serie_episode_servers
WHERE id = $1;