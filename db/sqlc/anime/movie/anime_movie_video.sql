-- name: CreateAnimeMovieVideo :one
INSERT INTO anime_movie_videos (language_id, autority, referer, link, quality)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAnimeMovieVideo :one
SELECT * FROM anime_movie_videos
WHERE id = $1
LIMIT 1;

-- name: UpdateAnimeMovieVideo :one
UPDATE anime_movie_videos
SET
    language_id = COALESCE(sqlc.narg(language_id), language_id),
    autority = COALESCE(sqlc.narg(autority), autority),
    referer = COALESCE(sqlc.narg(referer), referer),
    link = COALESCE(sqlc.narg(link), link),
    quality = COALESCE(sqlc.narg(quality), quality)
WHERE id = $1
RETURNING *;

-- name: ListAnimeMovieVideos :many
SELECT * FROM anime_movie_videos
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteAnimeMovieVideo :exec
DELETE FROM anime_movie_videos
WHERE id = $1;