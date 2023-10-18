ALTER TABLE "anime_movie_studio" DROP CONSTRAINT IF EXISTS "anime_movie_studio_studio_id_fkey";
ALTER TABLE "anime_movie_studio" DROP CONSTRAINT IF EXISTS "anime_movie_studio_anime_id_fkey";

ALTER TABLE "anime_movie_genre" DROP CONSTRAINT IF EXISTS "anime_movie_genre_genre_id_fkey";
ALTER TABLE "anime_movie_genre" DROP CONSTRAINT IF EXISTS "anime_movie_genre_anime_id_fkey";

ALTER TABLE "anime_movie_metas" DROP CONSTRAINT IF EXISTS "anime_movie_metas_anime_id_fkey";
ALTER TABLE "anime_movie_metas" DROP CONSTRAINT IF EXISTS "anime_movie_metas_language_id_fkey";
ALTER TABLE "anime_movie_metas" DROP CONSTRAINT IF EXISTS "anime_movie_metas_meta_id_fkey";

DROP TABLE IF EXISTS "anime_movie_metas";
DROP TABLE IF EXISTS "anime_movie_genre";
DROP TABLE IF EXISTS "anime_movie_studio";
DROP TABLE IF EXISTS "anime_movies";