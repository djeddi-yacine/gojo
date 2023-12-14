-- name: CreateAnimeSerieTrailer :one
INSERT INTO anime_serie_trailers (anime_id, trailer_id)
VALUES ($1, $2)
RETURNING *;

-- name: ListAnimeSerieTrailers :many
SELECT * FROM anime_serie_trailers
WHERE anime_id = $1;

-- name: DeleteAnimeSerieTrailer :exec
DELETE FROM anime_serie_trailers
WHERE anime_id = $1 AND trailer_id = $2;