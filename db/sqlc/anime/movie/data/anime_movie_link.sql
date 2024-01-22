-- name: CreateAnimeMovieLink :one
INSERT INTO anime_movie_links (anime_id, link_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieLink :one
SELECT * FROM anime_movie_links
WHERE anime_id = $1
LIMIT 1;

-- name: DeleteAnimeMovieLink :exec
DELETE FROM anime_movie_links
WHERE anime_id = $1 AND link_id = $2;