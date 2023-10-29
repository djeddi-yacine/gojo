-- name: CreateAnimeSerieServerDubVideo :one
INSERT INTO anime_serie_server_dub_videos (server_id, video_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieServerDubVideo :one
SELECT * FROM anime_serie_server_dub_videos
WHERE id = $1
LIMIT 1;

-- name: ListAnimeSerieServerDubVideos :many
SELECT * FROM anime_serie_server_dub_videos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateAnimeSerieServerDubVideo :one
UPDATE anime_serie_server_dub_videos
SET
  server_id = COALESCE(sqlc.narg(server_id), server_id),
  video_id = COALESCE(sqlc.narg(video_id), video_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSerieServerDubVideo :exec
DELETE FROM anime_serie_server_dub_videos
WHERE id = $1;
