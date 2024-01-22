-- name: GetAnimeMovieCharactersDirectly :many
SELECT anime_characters.*
FROM anime_characters
JOIN anime_movie_characters ON anime_characters.id = anime_movie_characters.studio_id
WHERE anime_movie_characters.anime_id = $1;
