ALTER TABLE "anime_movie_studios" DROP CONSTRAINT IF EXISTS "anime_movie_studios_studio_id_fkey";
ALTER TABLE "anime_movie_studios" DROP CONSTRAINT IF EXISTS "anime_movie_studios_anime_id_fkey";

ALTER TABLE "anime_movie_genres" DROP CONSTRAINT IF EXISTS "anime_movie_genres_genre_id_fkey";
ALTER TABLE "anime_movie_genres" DROP CONSTRAINT IF EXISTS "anime_movie_genres_anime_id_fkey";

ALTER TABLE "anime_movie_metas" DROP CONSTRAINT IF EXISTS "anime_movie_metas_anime_id_fkey";
ALTER TABLE "anime_movie_metas" DROP CONSTRAINT IF EXISTS "anime_movie_metas_language_id_fkey";
ALTER TABLE "anime_movie_metas" DROP CONSTRAINT IF EXISTS "anime_movie_metas_meta_id_fkey";

DROP TABLE IF EXISTS "anime_movie_metas";
DROP TABLE IF EXISTS "anime_movie_genres";
DROP TABLE IF EXISTS "anime_movie_studios";
DROP TABLE IF EXISTS "anime_movies";