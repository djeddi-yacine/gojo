-- name: CreateAnimeEpisodeVideo :one
INSERT INTO anime_episode_videos (language_id, authority, referer, link, quality)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAnimeEpisodeVideo :one
SELECT * FROM anime_episode_videos
WHERE id = $1
LIMIT 1;

-- name: UpdateAnimeEpisodeVideo :one
UPDATE anime_episode_videos
SET
    language_id = COALESCE(sqlc.narg(language_id), language_id),
    authority = COALESCE(sqlc.narg(authority), authority),
    referer = COALESCE(sqlc.narg(referer), referer),
    link = COALESCE(sqlc.narg(link), link),
    quality = COALESCE(sqlc.narg(quality), quality)
WHERE id = $1
RETURNING *;

-- name: ListAnimeEpisodeVideos :many
SELECT * FROM anime_episode_videos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteAnimeEpisodeVideo :exec
DELETE FROM anime_episode_videos
WHERE id = $1;