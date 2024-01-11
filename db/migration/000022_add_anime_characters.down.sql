ALTER TABLE "anime_movie_characters" DROP CONSTRAINT IF EXISTS "anime_movie_characters_anime_id_fkey";

ALTER TABLE "anime_movie_characters" DROP CONSTRAINT IF EXISTS "anime_movie_characters_character_id_fkey";

ALTER TABLE "anime_season_characters" DROP CONSTRAINT IF EXISTS "anime_season_characters_season_id_fkey";

ALTER TABLE "anime_season_characters" DROP CONSTRAINT IF EXISTS "anime_season_characters_character_id_fkey";

ALTER TABLE "anime_character_actors" DROP CONSTRAINT IF EXISTS "anime_character_actors_actor_id_fkey";

ALTER TABLE "anime_character_actors" DROP CONSTRAINT IF EXISTS "anime_character_actors_character_id_fkey";



DROP TABLE IF EXISTS "anime_movie_characters";
DROP TABLE IF EXISTS "anime_season_characters";
DROP TABLE IF EXISTS "anime_character_actors";
DROP TABLE IF EXISTS "anime_characters";
