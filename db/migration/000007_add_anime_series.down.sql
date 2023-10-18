ALTER TABLE "anime_serie_studios" DROP CONSTRAINT IF EXISTS "anime_serie_studios_studio_id_fkey";
ALTER TABLE "anime_serie_studios" DROP CONSTRAINT IF EXISTS "anime_serie_studios_anime_id_fkey";

ALTER TABLE "anime_serie_genres" DROP CONSTRAINT IF EXISTS "anime_serie_genres_genre_id_fkey";
ALTER TABLE "anime_serie_genres" DROP CONSTRAINT IF EXISTS "anime_serie_genres_anime_id_fkey";

ALTER TABLE "anime_serie_metas" DROP CONSTRAINT IF EXISTS "anime_serie_metas_anime_id_fkey";
ALTER TABLE "anime_serie_metas" DROP CONSTRAINT IF EXISTS "anime_serie_metas_language_id_fkey";
ALTER TABLE "anime_serie_metas" DROP CONSTRAINT IF EXISTS "anime_serie_metas_meta_id_fkey";

DROP TABLE IF EXISTS "anime_serie_metas";
DROP TABLE IF EXISTS "anime_serie_genres";
DROP TABLE IF EXISTS "anime_serie_studios";
DROP TABLE IF EXISTS "anime_series";