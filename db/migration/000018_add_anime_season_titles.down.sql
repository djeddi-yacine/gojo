ALTER TABLE "anime_season_official_titles" DROP CONSTRAINT IF EXISTS "anime_season_official_titles_season_id_fkey";

ALTER TABLE "anime_movie_short_titles" DROP CONSTRAINT IF EXISTS "anime_movie_short_titles_season_id_fkey";

ALTER TABLE "anime_movie_translation_titles" DROP CONSTRAINT IF EXISTS "anime_movie_translation_titles_season_id_fkey";

ALTER TABLE "anime_movie_other_titles" DROP CONSTRAINT IF EXISTS "anime_movie_other_titles_season_id_fkey";

DROP TABLE IF EXISTS "anime_season_official_titles";
DROP TABLE IF EXISTS "anime_season_short_titles";
DROP TABLE IF EXISTS "anime_season_translation_titles";
DROP TABLE IF EXISTS "anime_season_other_titles";