-- name: CreateAnimeSerieServerTorrent :one
INSERT INTO anime_serie_server_torrents (server_id, torrent_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieServerTorrent :one
SELECT * FROM anime_serie_server_torrents
WHERE id = $1
LIMIT 1;

-- name: ListAnimeSerieServerTorrents :many
SELECT * FROM anime_serie_server_torrents
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAnimeSerieServerTorrent :one
UPDATE anime_serie_server_torrents
SET 
  server_id = COALESCE(sqlc.narg(server_id), server_id),
  torrent_id = COALESCE(sqlc.narg(torrent_id), torrent_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSerieServerTorrent :exec
DELETE FROM anime_serie_server_torrents
WHERE id = $1;