-- name: GetAnimeSerieTrailersDirectly :many
SELECT anime_trailers.*
FROM anime_trailers
JOIN anime_serie_trailers ON anime_trailers.id = anime_serie_trailers.trailer_id
WHERE anime_serie_trailers.anime_id = $1;
