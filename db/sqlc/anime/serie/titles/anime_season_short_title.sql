-- name: CreateAnimeSeasonShortTitle :one
INSERT INTO anime_season_short_titles (season_id, title_text)
VALUES ($1, $2)
RETURNING *;

-- name: QueryAnimeSeasonShortTitles :many
WITH search_documents AS (
  SELECT
    season_id,
    title_text,
    to_tsvector('pg_catalog.english', title_text) AS title_text_tsv
  FROM anime_season_short_titles
)
SELECT season_id
FROM search_documents
WHERE (
  $1::text IS NOT NULL AND $1::text <> '' AND
  (
    SELECT bool_and(
      to_tsvector('pg_catalog.english', lower(translate(title_text, '[:punct:]', ''))) 
      @@ plainto_tsquery(word)
    )
    FROM unnest(regexp_split_to_array($1::text, '\\W+')) AS word
  )
  OR title_text ILIKE '%' || $1::text || '%'
)
ORDER BY
  ts_rank(title_text_tsv, phraseto_tsquery($1::text)) DESC,
  similarity(title_text, $1::text) DESC;

-- name: GetAnimeSeasonShortTitles :many
SELECT * FROM anime_season_short_titles
WHERE season_id = $1;

-- name: DeleteAnimeSeasonShortTitle :exec
DELETE FROM anime_season_short_titles
WHERE id = $1;