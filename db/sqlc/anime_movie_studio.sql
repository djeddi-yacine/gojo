-- name: CreateAnimeMovieStudio :one
INSERT INTO anime_movie_studio (anime_id, studio_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeMovieStudio :one
SELECT * FROM anime_movie_studio
WHERE id = $1 LIMIT 1;

-- name: ListAnimeMovieStudios :many
SELECT studio_id
FROM anime_movie_studio
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeMovieStudio :exec
DELETE FROM anime_movie_studio
WHERE anime_id = $1 AND studio_id = $2;