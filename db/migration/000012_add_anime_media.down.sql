ALTER TABLE "anime_movie_media" DROP CONSTRAINT IF EXISTS "anime_movie_media_anime_id_fkey";

ALTER TABLE "anime_movie_media" DROP CONSTRAINT IF EXISTS "anime_movie_media_media_id_fkey";

ALTER TABLE "anime_serie_media" DROP CONSTRAINT IF EXISTS "anime_serie_media_anime_id_fkey";

ALTER TABLE "anime_serie_media" DROP CONSTRAINT IF EXISTS "anime_serie_media_media_id_fkey";

ALTER TABLE "anime_serie_season_media" DROP CONSTRAINT IF EXISTS "anime_serie_season_media_season_id_fkey";

ALTER TABLE "anime_serie_season_media" DROP CONSTRAINT IF EXISTS "anime_serie_season_media_media_id_fkey";


DROP TABLE IF EXISTS "anime_movie_media";
DROP TABLE IF EXISTS "anime_serie_media";
DROP TABLE IF EXISTS "anime_serie_season_media";
DROP TABLE IF EXISTS "anime_media";