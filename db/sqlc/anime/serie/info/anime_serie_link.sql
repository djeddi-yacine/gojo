-- name: CreateAnimeSerieLink :one
INSERT INTO anime_serie_links (anime_id, link_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieLink :one
SELECT * FROM anime_serie_links
WHERE id = $1;

-- name: ListAnimeSerieLinksByAnimeID :many
SELECT link_id
FROM anime_serie_links
WHERE anime_id = $1
ORDER BY id;

-- name: DeleteAnimeSerieLink :exec
DELETE FROM anime_serie_links
WHERE anime_id = $1 AND link_id = $2;