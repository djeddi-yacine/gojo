-- name: CreateAnimeStudio :exec
INSERT INTO anime_studio (anime_id, studio_id)
VALUES ($1, $2);

-- name: GetAnimeStudio :one
SELECT * FROM anime_studio
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeStudio :one
UPDATE anime_studio
SET studio_id = $2
WHERE anime_id = $1
RETURNING * ;

-- name: ListAnimeStudios :many
SELECT studio_id
FROM anime_studio
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeStudio :exec
DELETE FROM anime_studio
WHERE anime_id = $1 AND studio_id = $2;