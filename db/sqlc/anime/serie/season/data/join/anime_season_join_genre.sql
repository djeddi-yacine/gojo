-- name: GetAnimeSeasonGenresDirectly :many
SELECT genres.*
FROM genres
JOIN anime_season_genres ON genres.id = anime_season_genres.genre_id
WHERE anime_season_genres.season_id = $1;
