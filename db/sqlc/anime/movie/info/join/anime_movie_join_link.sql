-- name: GetAnimeMovieLinksDirectly :one
SELECT anime_links.*
FROM anime_links
JOIN anime_movie_links ON anime_links.id = anime_movie_links.link_id
WHERE anime_movie_links.anime_id = $1
LIMIT 1;