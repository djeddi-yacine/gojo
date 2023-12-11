-- name: CreateAnimeSerieLink :one
INSERT INTO anime_serie_links (anime_id, link_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieLink :one
SELECT * FROM anime_serie_links
WHERE anime_id = $1
LIMIT 1;

-- name: DeleteAnimeSerieLink :exec
DELETE FROM anime_serie_links
WHERE anime_id = $1 AND link_id = $2;