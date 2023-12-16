-- name: CreateAnimeMovieOfficialTitle :one
INSERT INTO anime_movie_official_titles (anime_id, title_text)
VALUES ($1, $2)
RETURNING *;

-- name: QueryAnimeMovieOfficialTitles :many
WITH search_documents AS (
  SELECT
    anime_id,
    title_text,
    to_tsvector('pg_catalog.english', title_text) AS title_text_tsv
  FROM anime_movie_official_titles
)
SELECT anime_id
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

-- name: GetAnimeMovieOfficialTitles :many
SELECT * FROM anime_movie_official_titles
WHERE anime_id = $1;

-- name: DeleteAnimeMovieOfficialTitle :exec
DELETE FROM anime_movie_official_titles
WHERE id = $1;