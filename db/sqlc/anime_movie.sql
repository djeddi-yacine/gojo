-- name: GetAnimeMovie :one
SELECT * FROM anime_movie 
WHERE id = $1;

-- name: CreateAnimeMovie :one
INSERT INTO anime_movie (
    original_title,
    aired,
    release_year,
    duration
)
VALUES ($1, $2, $3, $4)
RETURNING *;

-- name: DeleteAnimeMovie :exec
DELETE FROM anime_movie 
WHERE id = $1;
