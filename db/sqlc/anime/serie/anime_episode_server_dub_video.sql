-- name: CreateAnimeEpisodeServerDubVideo :one
INSERT INTO anime_episode_server_dub_videos (server_id, video_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeEpisodeServerDubVideo :one
SELECT * FROM anime_episode_server_dub_videos
WHERE id = $1
LIMIT 1;

-- name: ListAnimeEpisodeServerDubVideos :many
SELECT * FROM anime_episode_server_dub_videos
WHERE server_id = $1
ORDER BY id;

-- name: UpdateAnimeEpisodeServerDubVideo :one
UPDATE anime_episode_server_dub_videos
SET
  server_id = COALESCE(sqlc.narg(server_id), server_id),
  video_id = COALESCE(sqlc.narg(video_id), video_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeEpisodeServerDubVideo :exec
DELETE FROM anime_episode_server_dub_videos
WHERE id = $1;
