ALTER TABLE "anime_movie_characters" DROP CONSTRAINT IF EXISTS "anime_movie_characters_anime_id_fkey";

ALTER TABLE "anime_movie_characters" DROP CONSTRAINT IF EXISTS "anime_movie_characters_character_id_fkey";

ALTER TABLE "anime_serie_characters" DROP CONSTRAINT IF EXISTS "anime_serie_characters_anime_id_fkey";

ALTER TABLE "anime_serie_characters" DROP CONSTRAINT IF EXISTS "anime_serie_characters_character_id_fkey";

ALTER TABLE "anime_characters" DROP CONSTRAINT IF EXISTS "anime_characters_actor_id_fkey";



DROP TABLE IF EXISTS "anime_movie_characters";
DROP TABLE IF EXISTS "anime_serie_characters";
DROP TABLE IF EXISTS "anime_characters";
