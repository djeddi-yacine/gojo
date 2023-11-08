-- name: CreateAnimeMovieMedia :one
INSERT INTO anime_movie_media (anime_id, media_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieMedia :one
SELECT * FROM anime_movie_media
WHERE id = $1 LIMIT 1;

-- name: GetAnimeMovieMediaByAnimeID :one
SELECT * FROM anime_movie_media
WHERE anime_id = $1 LIMIT 1;

-- name: ListAnimeMovieMedias :many
SELECT media_id
FROM anime_movie_media
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeMovieMedia :exec
DELETE FROM anime_movie_media
WHERE anime_id = $1 AND media_id = $2;