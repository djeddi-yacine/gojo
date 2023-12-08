-- name: CreateAnimeLink :one
INSERT INTO anime_links (official_website, wikipedia_url, crunchyroll_url, social_media)
VALUES ($1, $2, $3, $4)
RETURNING  *;

-- name: GetAnimeLink :one
SELECT * FROM anime_links
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeLink :one
UPDATE anime_links
SET
  official_website = COALESCE(sqlc.narg(official_website), official_website),
  wikipedia_url = COALESCE(sqlc.narg(wikipedia_url), wikipedia_url),
  crunchyroll_url = COALESCE(sqlc.narg(crunchyroll_url), crunchyroll_url),
  social_media = COALESCE(sqlc.narg(social_media), social_media)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeLink :exec
DELETE FROM anime_links
WHERE id = $1;