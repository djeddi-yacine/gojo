-- name: CreateAnimeStudio :exec
INSERT INTO anime_studios (anime_id, studio_id)
VALUES ($1, $2);

-- name: UpdateAnimeStudio :one
UPDATE anime_studios
SET studio_id = $2
WHERE anime_id = $1
RETURNING * ;

-- name: ListAnimeStudios :many
SELECT studio_id
FROM anime_studios
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeStudio :exec
DELETE FROM anime_studios
WHERE anime_id = $1 AND studio_id = $2;