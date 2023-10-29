-- name: CreateAnimeSerieSeasonEpisode :one
INSERT INTO anime_serie_season_episodes (season_id, episode_id)
VALUES ($1, $2)
RETURNING *;

-- name: GetAnimeSerieSeasonEpisode :one
SELECT * FROM anime_serie_season_episodes
WHERE id = $1
LIMIT 1;

-- name: ListAnimeSerieSeasonEpisodesBySeason :many
SELECT * FROM anime_serie_season_episodes
WHERE season_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;

-- name: UpdateAnimeSerieSeasonEpisode :one
UPDATE anime_serie_season_episodes
SET
  episode_id = COALESCE(sqlc.narg(episode_id), episode_id),
  season_id = COALESCE(sqlc.narg(season_id), season_id)
WHERE
  id = sqlc.arg(id)
RETURNING *;

-- name: DeleteAnimeSerieSeasonEpisode :exec
DELETE FROM anime_serie_season_episodes
WHERE id = $1;