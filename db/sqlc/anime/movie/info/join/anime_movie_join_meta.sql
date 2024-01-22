-- name: GetAnimeMovieMetaWithLanguageDirectly :one
SELECT m.*
FROM anime_movie_metas AS amm
JOIN metas AS m ON amm.meta_id = m.id
WHERE amm.anime_id = $1 AND amm.language_id = $2
LIMIT 1;