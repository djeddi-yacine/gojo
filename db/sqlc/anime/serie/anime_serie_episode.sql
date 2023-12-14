-- name: CreateAnimeEpisode :one
INSERT INTO anime_serie_episodes (
  episode_number,
  season_id,
  thumbnails,
  thumbnails_blur_hash
)
VALUES ($1, $2, $3, $4)
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
  thumbnails = COALESCE(sqlc.narg(thumbnails), thumbnails),
  episode_number = COALESCE(sqlc.narg(episode_number), episode_number),
  thumbnails_blur_hash = COALESCE(sqlc.narg(thumbnails_blur_hash), thumbnails_blur_hash)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeEpisode :exec
DELETE FROM anime_serie_episodes
WHERE id = $1;
