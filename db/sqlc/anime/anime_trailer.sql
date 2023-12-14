-- name: CreateAnimeTrailer :one
INSERT INTO anime_trailers (is_official, host_name, host_key)
VALUES ($1, $2, $3)
RETURNING  *;

-- name: GetAnimeTrailer :one
SELECT * FROM anime_trailers
WHERE id = $1 LIMIT 1;

-- name: UpdateAnimeTrailer :one
UPDATE anime_trailers
SET
  is_official = COALESCE(sqlc.narg(is_official), is_official),
  host_name = COALESCE(sqlc.narg(host_name), host_name),
  host_key = COALESCE(sqlc.narg(host_key), host_key)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeTrailer :exec
DELETE FROM anime_trailers
WHERE id = $1;