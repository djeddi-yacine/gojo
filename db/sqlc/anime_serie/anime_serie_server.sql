-- name: CreateAnimeSerieServer :one
INSERT INTO anime_serie_servers (episode_id)
VALUES ($1)
RETURNING *;

-- name: GetAnimeSerieServer :one
SELECT * FROM anime_serie_servers
WHERE id = $1
LIMIT 1;

-- name: UpdateAnimeSerieServer :one
UPDATE anime_serie_servers
SET episode_id = $2
WHERE id = $1
RETURNING *;

-- name: ListAnimeSerieServers :many
SELECT * FROM anime_serie_servers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteAnimeSerieServer :exec
DELETE FROM anime_serie_servers
WHERE id = $1;
