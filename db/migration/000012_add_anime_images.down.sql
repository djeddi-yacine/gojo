ALTER TABLE "anime_movie_poster_images" DROP CONSTRAINT IF EXISTS "anime_movie_poster_images_anime_id_fkey";

ALTER TABLE "anime_movie_backdrop_images" DROP CONSTRAINT IF EXISTS "anime_movie_backdrop_images_anime_id_fkey";

ALTER TABLE "anime_movie_logo_images" DROP CONSTRAINT IF EXISTS "anime_movie_logo_images_anime_id_fkey";

ALTER TABLE "anime_movie_poster_images" DROP CONSTRAINT IF EXISTS "anime_movie_poster_images_image_id_fkey";

ALTER TABLE "anime_movie_backdrop_images" DROP CONSTRAINT IF EXISTS "anime_movie_backdrop_images_image_id_fkey";

ALTER TABLE "anime_movie_logo_images" DROP CONSTRAINT IF EXISTS "anime_movie_logo_images_image_id_fkey";

ALTER TABLE "anime_serie_poster_images" DROP CONSTRAINT IF EXISTS "anime_serie_poster_images_anime_id_fkey";

ALTER TABLE "anime_serie_backdrop_images" DROP CONSTRAINT IF EXISTS "anime_serie_backdrop_images_anime_id_fkey";

ALTER TABLE "anime_serie_logo_images" DROP CONSTRAINT IF EXISTS "anime_serie_logo_images_anime_id_fkey";

ALTER TABLE "anime_serie_poster_images" DROP CONSTRAINT IF EXISTS "anime_serie_poster_images_image_id_fkey";

ALTER TABLE "anime_serie_backdrop_images" DROP CONSTRAINT IF EXISTS "anime_serie_backdrop_images_image_id_fkey";

ALTER TABLE "anime_serie_logo_images" DROP CONSTRAINT IF EXISTS "anime_serie_logo_images_image_id_fkey";

ALTER TABLE "anime_serie_season_poster_images" DROP CONSTRAINT IF EXISTS "anime_serie_season_poster_images_season_id_fkey";

ALTER TABLE "anime_serie_season_poster_images" DROP CONSTRAINT IF EXISTS "anime_serie_season_poster_images_images_id_fkey";


DROP TABLE IF EXISTS "anime_movie_poster_images";
DROP TABLE IF EXISTS "anime_movie_backdrop_images";
DROP TABLE IF EXISTS "anime_movie_logo_images";
DROP TABLE IF EXISTS "anime_serie_poster_images";
DROP TABLE IF EXISTS "anime_serie_backdrop_images";
DROP TABLE IF EXISTS "anime_serie_logo_images";
DROP TABLE IF EXISTS "anime_serie_season_poster_images";
DROP TABLE IF EXISTS "anime_images";