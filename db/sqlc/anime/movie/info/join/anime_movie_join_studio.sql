-- name: GetAnimeMovieStudiosDirectly :many
SELECT studios.*
FROM studios
JOIN anime_movie_studios ON studios.id = anime_movie_studios.studio_id
WHERE anime_movie_studios.anime_id = $1;
