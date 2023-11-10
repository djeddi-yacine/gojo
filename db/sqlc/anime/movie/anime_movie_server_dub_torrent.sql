-- name: CreateAnimeMovieServerDubTorrent :one
INSERT INTO anime_movie_server_dub_torrents (server_id, torrent_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieServerDubTorrent :one
SELECT * FROM anime_movie_server_dub_torrents
WHERE id = $1
LIMIT 1;

-- name: ListAnimeMovieServerDubTorrents :many
SELECT * FROM anime_movie_server_dub_torrents
WHERE server_id = $1
ORDER BY id;

-- name: UpdateAnimeMovieServerDubTorrent :one
UPDATE anime_movie_server_dub_torrents
SET 
  server_id = COALESCE(sqlc.narg(server_id), server_id),
  torrent_id = COALESCE(sqlc.narg(torrent_id), torrent_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeMovieServerDubTorrent :exec
DELETE FROM anime_movie_server_dub_torrents
WHERE id = $1;