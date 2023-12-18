-- name: CreateAnimeMovieTag :one
INSERT INTO anime_movie_tags (anime_id, tag_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListAnimeMovieTags :many
SELECT * FROM anime_movie_tags
WHERE anime_id = $1;

-- name: DeleteAnimeMovieTag :exec
DELETE FROM anime_movie_tags
WHERE anime_id = $1 AND tag_id = $2;