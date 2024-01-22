-- name: GetAnimeMovieTagsDirectly :many
SELECT anime_tags.*
FROM anime_tags
JOIN anime_movie_tags ON anime_tags.id = anime_movie_tags.tag_id
WHERE anime_movie_tags.anime_id = $1;
