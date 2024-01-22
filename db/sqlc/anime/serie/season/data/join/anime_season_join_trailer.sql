-- name: GetAnimeSeasonTrailersDirectly :many
SELECT anime_trailers.*
FROM anime_trailers
JOIN anime_season_trailers ON anime_trailers.id = anime_season_trailers.trailer_id
WHERE anime_season_trailers.season_id = $1;
