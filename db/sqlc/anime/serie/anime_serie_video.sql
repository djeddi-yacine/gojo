-- name: CreateAnimeSerieVideo :one
INSERT INTO anime_serie_videos (language_id, authority, referer, link, quality)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAnimeSerieVideo :one
SELECT * FROM anime_serie_videos
WHERE id = $1
LIMIT 1;

-- name: UpdateAnimeSerieVideo :one
UPDATE anime_serie_videos
SET
    language_id = COALESCE(sqlc.narg(language_id), language_id),
    authority = COALESCE(sqlc.narg(authority), authority),
    referer = COALESCE(sqlc.narg(referer), referer),
    link = COALESCE(sqlc.narg(link), link),
    quality = COALESCE(sqlc.narg(quality), quality)
WHERE id = $1
RETURNING *;

-- name: ListAnimeSerieVideos :many
SELECT * FROM anime_serie_videos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteAnimeSerieVideo :exec
DELETE FROM anime_serie_videos
WHERE id = $1;