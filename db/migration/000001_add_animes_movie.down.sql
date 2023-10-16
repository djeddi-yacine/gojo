ALTER TABLE "anime_studio" DROP CONSTRAINT IF EXISTS "anime_studio_studio_id_fkey";
ALTER TABLE "anime_genre" DROP CONSTRAINT IF EXISTS "anime_genre_genre_id_fkey";
ALTER TABLE "anime_metas" DROP CONSTRAINT IF EXISTS "anime_metas_anime_id_fkey";
ALTER TABLE "anime_studio" DROP CONSTRAINT IF EXISTS "anime_studio_anime_id_fkey";
ALTER TABLE "anime_genre" DROP CONSTRAINT IF EXISTS "anime_genre_anime_id_fkey";
ALTER TABLE "anime_metas" DROP CONSTRAINT IF EXISTS "anime_metas_language_id_fkey";
ALTER TABLE "anime_metas" DROP CONSTRAINT IF EXISTS "anime_metas_meta_id_fkey";

DROP TABLE IF EXISTS "anime_metas";
DROP TABLE IF EXISTS "metas";
DROP TABLE IF EXISTS "languages";
DROP TABLE IF EXISTS "anime_genre";
DROP TABLE IF EXISTS "genres";
DROP TABLE IF EXISTS "anime_studio";
DROP TABLE IF EXISTS "studios";
DROP TABLE IF EXISTS "anime_movie";