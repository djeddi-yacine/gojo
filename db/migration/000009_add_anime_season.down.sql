ALTER TABLE "anime_season_metas" DROP CONSTRAINT IF EXISTS "anime_season_metas_meta_id_fkey";

ALTER TABLE "anime_season_metas" DROP CONSTRAINT IF EXISTS "anime_season_metas_season_id_fkey";

ALTER TABLE "anime_season_metas" DROP CONSTRAINT IF EXISTS "anime_season_metas_language_id_fkey";

ALTER TABLE "anime_serie_seasons" DROP CONSTRAINT IF EXISTS "anime_serie_seasons_anime_id_fkey";

ALTER TABLE "anime_season_studios" DROP CONSTRAINT IF EXISTS "anime_season_studios_studio_id_fkey";

ALTER TABLE "anime_season_studios" DROP CONSTRAINT IF EXISTS "anime_season_studios_anime_id_fkey";

ALTER TABLE "anime_season_genres" DROP CONSTRAINT IF EXISTS "anime_season_genres_genre_id_fkey";

ALTER TABLE "anime_season_genres" DROP CONSTRAINT IF EXISTS "anime_season_genres_anime_id_fkey";


DROP TABLE IF EXISTS "anime_season_genres";
DROP TABLE IF EXISTS "anime_season_studios";
DROP TABLE IF EXISTS "anime_season_metas";
DROP TABLE IF EXISTS "anime_serie_seasons";