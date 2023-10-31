-- name: CreateAnimeResource :one
INSERT INTO anime_resources (tmdb_id, imdb_id, official_website, wikipedia_url, crunchyroll_url, social_media)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING  *;

-- name: GetAnimeResource :one
SELECT * FROM anime_resources
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeResource :one
UPDATE anime_resources
SET
  tmdb_id = COALESCE(sqlc.narg(tmdb_id), tmdb_id),
  imdb_id = COALESCE(sqlc.narg(imdb_id), imdb_id),
  official_website = COALESCE(sqlc.narg(official_website), official_website),
  wikipedia_url = COALESCE(sqlc.narg(wikipedia_url), wikipedia_url),
  crunchyroll_url = COALESCE(sqlc.narg(crunchyroll_url), crunchyroll_url)
  social_media = COALESCE(sqlc.narg(social_media), social_media)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeResource :exec
DELETE FROM anime_resources
WHERE id = $1;