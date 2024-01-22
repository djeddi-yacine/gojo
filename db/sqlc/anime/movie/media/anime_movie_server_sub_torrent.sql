-- name: CreateAnimeMovieServerSubTorrent :one
INSERT INTO anime_movie_server_sub_torrents (server_id, torrent_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieServerSubTorrent :one
SELECT * FROM anime_movie_server_sub_torrents
WHERE id = $1
LIMIT 1;

-- name: ListAnimeMovieServerSubTorrents :many
SELECT * FROM anime_movie_server_sub_torrents
WHERE server_id = $1
ORDER BY id;

-- name: UpdateAnimeMovieServerSubTorrent :one
UPDATE anime_movie_server_sub_torrents
SET 
  server_id = COALESCE(sqlc.narg(server_id), server_id),
  torrent_id = COALESCE(sqlc.narg(torrent_id), torrent_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeMovieServerSubTorrent :exec
DELETE FROM anime_movie_server_sub_torrents
WHERE id = $1;