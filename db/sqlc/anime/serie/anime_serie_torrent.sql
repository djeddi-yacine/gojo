-- name: CreateAnimeSerieTorrent :one
INSERT INTO anime_serie_torrents (language_id, file_name, torrent_hash, torrent_file, seeds, peers, leechers, size_bytes, quality)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
RETURNING *;

-- name: GetAnimeSerieTorrent :one
SELECT * FROM anime_serie_torrents
WHERE id = $1
LIMIT 1;

-- name: UpdateAnimeSerieTorrent :one
UPDATE anime_serie_torrents
SET
    language_id = COALESCE(sqlc.narg(language_id), language_id),
    file_name = COALESCE(sqlc.narg(file_name), file_name),
    torrent_hash = COALESCE(sqlc.narg(torrent_hash), torrent_hash),
    torrent_file = COALESCE(sqlc.narg(torrent_file), torrent_file),
    seeds = COALESCE(sqlc.narg(seeds), seeds),
    peers = COALESCE(sqlc.narg(peers), peers),
    leechers = COALESCE(sqlc.narg(leechers), leechers),
    size_bytes = COALESCE(sqlc.narg(size_bytes), size_bytes),
    quality = COALESCE(sqlc.narg(quality), quality)
WHERE id = $1
RETURNING *;

-- name: ListAnimeSerieTorrents :many
SELECT * FROM anime_serie_torrents
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteAnimeSerieTorrent :exec
DELETE FROM anime_serie_torrents
WHERE id = $1;