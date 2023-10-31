ALTER TABLE "anime_movie_resources" DROP CONSTRAINT IF EXISTS "anime_movie_resources_anime_id_fkey";

ALTER TABLE "anime_movie_resources" DROP CONSTRAINT IF EXISTS "anime_movie_resources_resource_id_fkey";

ALTER TABLE "anime_serie_resources" DROP CONSTRAINT IF EXISTS "anime_serie_resources_anime_id_fkey";

ALTER TABLE "anime_serie_resources" DROP CONSTRAINT IF EXISTS "anime_serie_resources_resource_id_fkey";

DROP TABLE IF EXISTS "anime_movie_resources";
DROP TABLE IF EXISTS "anime_serie_resources";
DROP TABLE IF EXISTS "anime_resources";