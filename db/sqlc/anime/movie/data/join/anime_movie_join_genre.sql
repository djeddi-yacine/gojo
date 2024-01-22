-- name: GetAnimeMovieGenresDirectly :many
SELECT genres.*
FROM genres
JOIN anime_movie_genres ON genres.id = anime_movie_genres.genre_id
WHERE anime_movie_genres.anime_id = $1;
