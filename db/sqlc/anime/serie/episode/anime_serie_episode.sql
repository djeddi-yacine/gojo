-- name: CreateAnimeEpisode :one
INSERT INTO anime_serie_episodes (
  season_id,
  episode_number,
  episode_original_title,
  aired,
  rating,
  duration,
  thumbnails,
  thumbnails_blur_hash
)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetAnimeEpisodeByEpisodeID :one
SELECT * FROM anime_serie_episodes
WHERE id = $1
LIMIT 1;

-- name: ListAnimeEpisodesBySeasonID :many
SELECT * FROM anime_serie_episodes
WHERE season_id = $1
ORDER BY episode_number
LIMIT $2
OFFSET $3;

-- name: UpdateAnimeEpisode :one
UPDATE anime_serie_episodes
SET
  episode_number = COALESCE(sqlc.narg(episode_number), episode_number),
  episode_original_title = COALESCE(sqlc.narg(episode_original_title), episode_original_title),
  aired = COALESCE(sqlc.narg(aired), aired),
  rating = COALESCE(sqlc.narg(rating), rating),
  duration = COALESCE(sqlc.narg(duration), duration),
  thumbnails = COALESCE(sqlc.narg(thumbnails), thumbnails),
  thumbnails_blur_hash = COALESCE(sqlc.narg(thumbnails_blur_hash), thumbnails_blur_hash)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeEpisode :exec
DELETE FROM anime_serie_episodes
WHERE id = $1;
