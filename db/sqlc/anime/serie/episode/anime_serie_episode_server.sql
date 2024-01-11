-- name: CreateAnimeEpisodeServer :one
INSERT INTO anime_episode_servers (episode_id)
VALUES ($1)
RETURNING *;

-- name: GetAnimeEpisodeServer :one
SELECT * FROM anime_episode_servers
WHERE id = $1
LIMIT 1;

-- name: GetAnimeEpisodeServerByEpisodeID :one
SELECT id FROM anime_episode_servers
WHERE episode_id = $1
LIMIT 1;

-- name: UpdateAnimeEpisodeServer :one
UPDATE anime_episode_servers
SET episode_id = $2
WHERE id = $1
RETURNING *;

-- name: ListAnimeEpisodeServers :many
SELECT * FROM anime_episode_servers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteAnimeEpisodeServer :exec
DELETE FROM anime_episode_servers
WHERE id = $1;
