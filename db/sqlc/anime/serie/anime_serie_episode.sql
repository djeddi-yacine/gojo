-- name: CreateAnimeSerieEpisode :one
INSERT INTO anime_serie_episodes (episode_number, season_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieEpisode :one
SELECT * FROM anime_serie_episodes
WHERE id = $1
LIMIT 1;

-- name: ListAnimeSerieEpisodesBySeason :many
SELECT * FROM anime_serie_episodes
WHERE season_id = $1
ORDER BY episode_number
LIMIT $2
OFFSET $3;

-- name: UpdateAnimeSerieEpisode :one
UPDATE anime_serie_episodes
SET
  episode_number = COALESCE(sqlc.narg(episode_number), episode_number),
  season_id = COALESCE(sqlc.narg(season_id), season_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSerieEpisode :exec
DELETE FROM anime_serie_episodes
WHERE id = $1;
