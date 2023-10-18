ALTER TABLE "anime_serie_studio" DROP CONSTRAINT IF EXISTS "anime_serie_studio_studio_id_fkey";
ALTER TABLE "anime_serie_studio" DROP CONSTRAINT IF EXISTS "anime_serie_studio_anime_id_fkey";

ALTER TABLE "anime_serie_genre" DROP CONSTRAINT IF EXISTS "anime_serie_genre_genre_id_fkey";
ALTER TABLE "anime_serie_genre" DROP CONSTRAINT IF EXISTS "anime_serie_genre_anime_id_fkey";

ALTER TABLE "anime_serie_metas" DROP CONSTRAINT IF EXISTS "anime_serie_metas_anime_id_fkey";
ALTER TABLE "anime_serie_metas" DROP CONSTRAINT IF EXISTS "anime_serie_metas_language_id_fkey";
ALTER TABLE "anime_serie_metas" DROP CONSTRAINT IF EXISTS "anime_serie_metas_meta_id_fkey";

DROP TABLE IF EXISTS "anime_serie_metas";
DROP TABLE IF EXISTS "anime_serie_genre";
DROP TABLE IF EXISTS "anime_serie_studio";
DROP TABLE IF EXISTS "anime_series";