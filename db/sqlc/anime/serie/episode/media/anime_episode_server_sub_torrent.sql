-- name: CreateAnimeEpisodeServerSubTorrent :one
INSERT INTO anime_episode_server_sub_torrents (server_id, torrent_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeEpisodeServerSubTorrent :one
SELECT * FROM anime_episode_server_sub_torrents
WHERE id = $1
LIMIT 1;

-- name: ListAnimeEpisodeServerSubTorrents :many
SELECT * FROM anime_episode_server_sub_torrents
WHERE server_id = $1
ORDER BY id;

-- name: UpdateAnimeEpisodeServerSubTorrent :one
UPDATE anime_episode_server_sub_torrents
SET 
  server_id = COALESCE(sqlc.narg(server_id), server_id),
  torrent_id = COALESCE(sqlc.narg(torrent_id), torrent_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeEpisodeServerSubTorrent :exec
DELETE FROM anime_episode_server_sub_torrents
WHERE id = $1;