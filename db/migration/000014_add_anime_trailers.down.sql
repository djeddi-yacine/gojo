ALTER TABLE "anime_movie_trailers" DROP CONSTRAINT IF EXISTS "anime_movie_trailers_anime_id_fkey";

ALTER TABLE "anime_movie_trailers" DROP CONSTRAINT IF EXISTS "anime_movie_trailers_trailer_id_fkey";

ALTER TABLE "anime_serie_trailers" DROP CONSTRAINT IF EXISTS "anime_serie_trailers_anime_id_fkey";

ALTER TABLE "anime_serie_trailers" DROP CONSTRAINT IF EXISTS "anime_serie_trailers_trailer_id_fkey";

ALTER TABLE "anime_season_trailers" DROP CONSTRAINT IF EXISTS "anime_season_trailers_anime_id_fkey";

ALTER TABLE "anime_season_trailers" DROP CONSTRAINT IF EXISTS "anime_season_trailers_trailer_id_fkey";

DROP TABLE IF EXISTS "anime_movie_trailers";
DROP TABLE IF EXISTS "anime_serie_trailers";
DROP TABLE IF EXISTS "anime_season_trailers";
DROP TABLE IF EXISTS "anime_trailers";