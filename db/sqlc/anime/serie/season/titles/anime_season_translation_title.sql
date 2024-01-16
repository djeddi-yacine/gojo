-- name: CreateAnimeSeasonTranslationTitle :one
INSERT INTO anime_season_translation_titles (season_id, title_text)
VALUES ($1, $2)
ON CONFLICT (season_id, title_text)
DO UPDATE SET 
    season_id = excluded.season_id
RETURNING *;

-- name: QueryAnimeSeasonTranslationTitles :many
WITH search_documents AS (
  SELECT
    season_id,
    title_text,
    to_tsvector('pg_catalog.simple', title_text) AS title_text_tsv
  FROM anime_season_translation_titles
)
SELECT season_id
FROM search_documents
WHERE (
  $1::varchar IS NOT NULL AND $1::varchar <> '' AND
  (
    SELECT bool_and(
      to_tsvector('pg_catalog.simple', lower(translate(title_text, '[:punct:]', ''))) 
      @@ plainto_tsquery(word)
    )
    FROM unnest(regexp_split_to_array($1::varchar, '\\W+')) AS word
  )
  OR title_text ILIKE '%' || $1::varchar || '%'
)
ORDER BY
  ts_rank(title_text_tsv, phraseto_tsquery($1::varchar)) DESC,
  similarity(title_text, $1::varchar) DESC
  LIMIT $2
  OFFSET $3;

-- name: GetAnimeSeasonTranslationTitles :many
SELECT * FROM anime_season_translation_titles
WHERE season_id = $1;

-- name: DeleteAnimeSeasonTranslationTitle :exec
DELETE FROM anime_season_translation_titles
WHERE id = $1;