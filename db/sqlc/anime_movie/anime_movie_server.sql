-- name: CreateAnimeMovieServer :one
INSERT INTO anime_movie_servers (anime_id)
VALUES ($1)
RETURNING *;

-- name: GetAnimeMovieServer :one
SELECT * FROM anime_movie_servers
WHERE id = $1
LIMIT 1;

-- name: UpdateAnimeMovieServer :one
UPDATE anime_movie_servers
SET anime_id = $2
WHERE id = $1
RETURNING *;

-- name: ListAnimeMovieServers :many
SELECT * FROM anime_movie_servers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: DeleteAnimeMovieServer :exec
DELETE FROM anime_movie_servers
WHERE id = $1;
