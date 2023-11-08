-- name: CreateAnimeMedia :one
INSERT INTO anime_media (media_type, media_url, media_author, media_platform)
VALUES ($1, $2, $3, $4)
RETURNING  *;

-- name: GetAnimeMedia :one
SELECT * FROM anime_media
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeMedia :one
UPDATE anime_media
SET
  media_type = COALESCE(sqlc.narg(media_type), media_type),
  media_url = COALESCE(sqlc.narg(media_url), media_url),
  media_author = COALESCE(sqlc.narg(media_author), media_author),
  media_platform = COALESCE(sqlc.narg(media_platform), media_platform)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeMedia :exec
DELETE FROM anime_media
WHERE id = $1;