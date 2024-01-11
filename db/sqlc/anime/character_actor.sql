-- name: CreateAnimeCharacterActor :exec
INSERT INTO anime_character_actors (character_id, actor_id)
VALUES ($1, $2)
ON CONFLICT (character_id, actor_id)
DO NOTHING;

-- name: ListAnimeCharacterActors :many
SELECT * FROM anime_character_actors
WHERE character_id = $1;

-- name: DeleteAnimeCharacterActor :exec
DELETE FROM anime_character_actors
WHERE id = $1;