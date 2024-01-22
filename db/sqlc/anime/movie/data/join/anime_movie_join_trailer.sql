-- name: GetAnimeMovieTrailersDirectly :many
SELECT anime_trailers.*
FROM anime_trailers
JOIN anime_movie_trailers ON anime_trailers.id = anime_movie_trailers.trailer_id
WHERE anime_movie_trailers.anime_id = $1;
