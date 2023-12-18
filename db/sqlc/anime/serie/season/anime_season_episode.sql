-- name: CreateAnimeSeasonEpisode :one
INSERT INTO anime_season_episodes (season_id, episode_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSeasonEpisode :one
SELECT * FROM anime_season_episodes
WHERE id = $1
LIMIT 1;

-- name: ListAnimeSeasonEpisodesBySeason :many
SELECT * FROM anime_season_episodes
WHERE season_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAnimeSeasonEpisode :one
UPDATE anime_season_episodes
SET
  episode_id = COALESCE(sqlc.narg(episode_id), episode_id),
  season_id = COALESCE(sqlc.narg(season_id), season_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSeasonEpisode :exec
DELETE FROM anime_season_episodes
WHERE id = $1;