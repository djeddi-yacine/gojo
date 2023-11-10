-- name: CreateAnimeSerieServerDubTorrent :one
INSERT INTO anime_serie_server_dub_torrents (server_id, torrent_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieServerDubTorrent :one
SELECT * FROM anime_serie_server_dub_torrents
WHERE id = $1
LIMIT 1;

-- name: ListAnimeSerieServerDubTorrents :many
SELECT * FROM anime_serie_server_dub_torrents
WHERE server_id = $1
ORDER BY id;

-- name: UpdateAnimeSerieServerDubTorrent :one
UPDATE anime_serie_server_dub_torrents
SET 
  server_id = COALESCE(sqlc.narg(server_id), server_id),
  torrent_id = COALESCE(sqlc.narg(torrent_id), torrent_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSerieServerDubTorrent :exec
DELETE FROM anime_serie_server_dub_torrents
WHERE id = $1;