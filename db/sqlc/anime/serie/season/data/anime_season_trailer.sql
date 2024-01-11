-- name: CreateAnimeSeasonTrailer :one
INSERT INTO anime_season_trailers (season_id, trailer_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListAnimeSeasonTrailers :many
SELECT trailer_id FROM anime_season_trailers
WHERE season_id = $1;

-- name: DeleteAnimeSeasonTrailer :exec
DELETE FROM anime_season_trailers
WHERE season_id = $1 AND trailer_id = $2;