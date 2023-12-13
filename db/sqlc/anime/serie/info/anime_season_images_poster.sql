-- name: CreateAnimeSeasonPosterImage :one
INSERT INTO anime_season_poster_images (season_id, image_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSeasonPosterImage :one
SELECT * FROM anime_season_poster_images
WHERE id = $1 LIMIT 1;

-- name: GetAnimeSeasonPosterImageByAnimeID :one
SELECT * FROM anime_season_poster_images
WHERE season_id = $1 LIMIT 1;

-- name: ListAnimeSeasonPosterImages :many
SELECT image_id
FROM anime_season_poster_images
WHERE season_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeSeasonPosterImage :exec
DELETE FROM anime_season_poster_images
WHERE season_id = $1 AND image_id = $2;