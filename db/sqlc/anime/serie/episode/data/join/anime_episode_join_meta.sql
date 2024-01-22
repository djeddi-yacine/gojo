-- name: GetAnimeEpisodeMetaWithLanguageDirectly :one
SELECT m.*
FROM anime_episode_metas AS aem
JOIN metas AS m ON aem.meta_id = m.id
WHERE aem.episode_id = $1 AND aem.language_id = $2
LIMIT 1;