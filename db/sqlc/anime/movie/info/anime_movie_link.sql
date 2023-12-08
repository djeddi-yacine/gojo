-- name: CreateAnimeMovieLink :one
INSERT INTO anime_movie_links (anime_id, link_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieLink :one
SELECT * FROM anime_movie_links
WHERE id = $1;

-- name: ListAnimeMovieLinksByAnimeID :many
SELECT * FROM anime_movie_links
WHERE anime_id = $1 
ORDER BY id;

-- name: DeleteAnimeMovieLink :exec
DELETE FROM anime_movie_links
WHERE anime_id = $1 AND link_id = $2;