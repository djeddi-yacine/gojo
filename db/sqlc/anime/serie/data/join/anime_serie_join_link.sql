-- name: GetAnimeSerieLinksDirectly :one
SELECT anime_links.*
FROM anime_links
JOIN anime_serie_links ON anime_links.id = anime_serie_links.link_id
WHERE anime_serie_links.anime_id = $1
LIMIT 1;