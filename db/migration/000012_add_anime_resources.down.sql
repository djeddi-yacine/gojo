ALTER TABLE "anime_movie_resources" DROP CONSTRAINT IF EXISTS "anime_movie_resources_anime_id_fkey";

ALTER TABLE "anime_movie_resources" DROP CONSTRAINT IF EXISTS "anime_movie_resources_resource_id_fkey";

ALTER TABLE "anime_season_resources" DROP CONSTRAINT IF EXISTS "anime_season_resources_season_id_fkey";

ALTER TABLE "anime_season_resources" DROP CONSTRAINT IF EXISTS "anime_season_resources_resource_id_fkey";

DROP TABLE IF EXISTS "anime_movie_resources";
DROP TABLE IF EXISTS "anime_season_resources";
DROP TABLE IF EXISTS "anime_resources";