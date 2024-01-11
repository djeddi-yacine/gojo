-- name: CreateAnimeMovieTrailer :one
INSERT INTO anime_movie_trailers (anime_id, trailer_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListAnimeMovieTrailers :many
SELECT trailer_id FROM anime_movie_trailers
WHERE anime_id = $1;

-- name: DeleteAnimeMovieTrailer :exec
DELETE FROM anime_movie_trailers
WHERE anime_id = $1 AND trailer_id = $2;