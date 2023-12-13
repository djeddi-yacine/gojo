-- name: CreateAnimeSerieSeasonPosterImage :one
INSERT INTO anime_serie_season_poster_images (season_id, image_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieSeasonPosterImage :one
SELECT * FROM anime_serie_season_poster_images
WHERE id = $1 LIMIT 1;

-- name: GetAnimeSerieSeasonPosterImageByAnimeID :one
SELECT * FROM anime_serie_season_poster_images
WHERE season_id = $1 LIMIT 1;

-- name: ListAnimeSerieSeasonPosterImages :many
SELECT image_id
FROM anime_serie_season_poster_images
WHERE season_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeSerieSeasonPosterImage :exec
DELETE FROM anime_serie_season_poster_images
WHERE season_id = $1 AND image_id = $2;