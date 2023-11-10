-- name: CreateAnimeMovieResource :one
INSERT INTO anime_movie_resources (anime_id, resource_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieResource :one
SELECT * FROM anime_movie_resources
WHERE id = $1;

-- name: ListAnimeMovieResourcesByAnimeID :many
SELECT * FROM anime_movie_resources
WHERE anime_id = $1 
ORDER BY id;

-- name: DeleteAnimeMovieResource :exec
DELETE FROM anime_movie_resources
WHERE anime_id = $1 AND resource_id = $2;