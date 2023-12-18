ALTER TABLE "anime_movie_tags" DROP CONSTRAINT IF EXISTS "anime_movie_tags_anime_id_fkey";

ALTER TABLE "anime_season_tags" DROP CONSTRAINT IF EXISTS "anime_season_tags_season_id_fkey";

ALTER TABLE "anime_movie_tags" DROP CONSTRAINT IF EXISTS "anime_movie_tags_tag_id_fkey";

ALTER TABLE "anime_season_tags" DROP CONSTRAINT IF EXISTS "anime_season_tags_tag_id_fkey";


DROP TABLE IF EXISTS "anime_season_tags";
DROP TABLE IF EXISTS "anime_movie_tags";
DROP TABLE IF EXISTS "anime_tags";