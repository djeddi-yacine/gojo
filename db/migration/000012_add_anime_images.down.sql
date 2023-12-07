ALTER TABLE "anime_movie_images" DROP CONSTRAINT IF EXISTS "anime_movie_images_anime_id_fkey";

ALTER TABLE "anime_movie_images" DROP CONSTRAINT IF EXISTS "anime_movie_images_image_id_fkey";

ALTER TABLE "anime_serie_images" DROP CONSTRAINT IF EXISTS "anime_serie_images_anime_id_fkey";

ALTER TABLE "anime_serie_images" DROP CONSTRAINT IF EXISTS "anime_serie_images_image_id_fkey";

ALTER TABLE "anime_serie_season_images" DROP CONSTRAINT IF EXISTS "anime_serie_season_images_season_id_fkey";

ALTER TABLE "anime_serie_season_images" DROP CONSTRAINT IF EXISTS "anime_serie_season_images_images_id_fkey";


DROP TABLE IF EXISTS "anime_movie_images";
DROP TABLE IF EXISTS "anime_serie_images";
DROP TABLE IF EXISTS "anime_serie_season_images";
DROP TABLE IF EXISTS "anime_images";