-- name: CreateAnimeSerieStudio :one
INSERT INTO anime_serie_studios (anime_id, studio_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieStudio :one
SELECT * FROM anime_serie_studios
WHERE anime_id = $1 AND studio_id = $2;

-- name: ListAnimeSerieStudios :many
SELECT studio_id
FROM anime_serie_studios
WHERE anime_id = $1
ORDER BY id;

-- name: DeleteAnimeSerieStudio :exec
DELETE FROM anime_serie_studios
WHERE anime_id = $1 AND studio_id = $2;