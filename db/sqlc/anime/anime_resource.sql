-- name: CreateAnimeResource :one
INSERT INTO anime_resources (
  tvdb_id,
  tmdb_id,
  imdb_id,
  livechart_id,
  anime_planet_id,
  anisearch_id,
  anidb_id,
  kitsu_id,
  mal_id,
  notify_moe_id,
  anilist_id
  ) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
RETURNING  *;

-- name: GetAnimeResource :one
SELECT * FROM anime_resources
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeResource :one
UPDATE anime_resources
SET
  tvdb_id = COALESCE(sqlc.narg(tvdb_id), tvdb_id),
  tmdb_id = COALESCE(sqlc.narg(tmdb_id), tmdb_id),
  imdb_id = COALESCE(sqlc.narg(imdb_id), imdb_id),
  livechart_id = COALESCE(sqlc.narg(livechart_id), livechart_id),
  anime_planet_id = COALESCE(sqlc.narg(anime_planet_id), anime_planet_id),
  anisearch_id = COALESCE(sqlc.narg(anisearch_id), anisearch_id),
  anidb_id = COALESCE(sqlc.narg(anidb_id), anidb_id),
  kitsu_id = COALESCE(sqlc.narg(kitsu_id), kitsu_id),
  mal_id = COALESCE(sqlc.narg(mal_id), mal_id),
  notify_moe_id = COALESCE(sqlc.narg(notify_moe_id), notify_moe_id),
  anilist_id = COALESCE(sqlc.narg(anilist_id), anilist_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeResource :exec
DELETE FROM anime_resources
WHERE id = $1;