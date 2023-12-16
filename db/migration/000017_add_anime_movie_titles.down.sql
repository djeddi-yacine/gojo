ALTER TABLE "anime_movie_official_titles" DROP CONSTRAINT IF EXISTS "anime_movie_official_titles_anime_id_fkey";

ALTER TABLE "anime_movie_short_titles" DROP CONSTRAINT IF EXISTS "anime_movie_short_titles_anime_id_fkey";

ALTER TABLE "anime_movie_translation_titles" DROP CONSTRAINT IF EXISTS "anime_movie_translation_titles_anime_id_fkey";

ALTER TABLE "anime_movie_other_titles" DROP CONSTRAINT IF EXISTS "anime_movie_other_titles_anime_id_fkey";

DROP TABLE IF EXISTS "anime_movie_official_titles";
DROP TABLE IF EXISTS "anime_movie_short_titles";
DROP TABLE IF EXISTS "anime_movie_translation_titles";
DROP TABLE IF EXISTS "anime_movie_other_titles";