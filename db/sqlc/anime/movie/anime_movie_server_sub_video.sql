-- name: CreateAnimeMovieServerSubVideo :one
INSERT INTO anime_movie_server_sub_videos (server_id, video_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieServerSubVideo :one
SELECT * FROM anime_movie_server_sub_videos
WHERE id = $1
LIMIT 1;

-- name: ListAnimeMovieServerSubVideos :many
SELECT * FROM anime_movie_server_sub_videos
WHERE server_id = $1
ORDER BY id;

-- name: UpdateAnimeMovieServerSubVideo :one
UPDATE anime_movie_server_sub_videos
SET 
  server_id = COALESCE(sqlc.narg(server_id), server_id),
  video_id = COALESCE(sqlc.narg(video_id), video_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeMovieServerSubVideo :exec
DELETE FROM anime_movie_server_sub_videos
WHERE id = $1;