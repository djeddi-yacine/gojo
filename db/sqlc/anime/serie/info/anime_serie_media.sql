-- name: CreateAnimeSerieMedia :one
INSERT INTO anime_serie_media (anime_id, media_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieMedia :one
SELECT * FROM anime_serie_media
WHERE id = $1 LIMIT 1;

-- name: GetAnimeSerieMediaByAnimeID :one
SELECT * FROM anime_serie_media
WHERE anime_id = $1 LIMIT 1;

-- name: ListAnimeSerieMedias :many
SELECT media_id
FROM anime_serie_media
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeSerieMedia :exec
DELETE FROM anime_serie_media
WHERE anime_id = $1 AND media_id = $2;