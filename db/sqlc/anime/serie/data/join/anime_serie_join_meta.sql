-- name: GetAnimeSerieMetaWithLanguageDirectly :one
SELECT m.*
FROM anime_serie_metas AS asm
JOIN metas AS m ON asm.meta_id = m.id
WHERE asm.anime_id = $1 AND asm.language_id = $2
LIMIT 1;