-- name: GetAnimeMovieResourceDirectly :one
SELECT anime_resources.*
FROM anime_resources
JOIN anime_movie_resources ON anime_resources.id = anime_movie_resources.resource_id
WHERE anime_movie_resources.anime_id = $1
LIMIT 1;