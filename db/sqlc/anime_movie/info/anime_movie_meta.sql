-- name: CreateAnimeMovieMeta :one
INSERT INTO anime_movie_metas (anime_id, language_id, meta_id)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetAnimeMovieMeta :one
SELECT meta_id
FROM anime_movie_metas
WHERE anime_id = $1 AND language_id = $2;

-- name: UpdateAnimeMovieMeta :one
UPDATE anime_movie_metas
SET meta_id = $3
WHERE anime_id = $1 AND language_id = $2
RETURNING * ;

-- name: ListAnimeMovieMetas :many
SELECT meta_id
FROM anime_movie_metas
WHERE anime_id = $1
LIMIT $2
OFFSET $3;

-- name: DeleteAnimeMovieMeta :exec
DELETE FROM anime_movie_metas
WHERE anime_id = $1 AND language_id = $2;