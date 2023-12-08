ALTER TABLE "anime_movie_links" DROP CONSTRAINT IF EXISTS "anime_movie_links_anime_id_fkey";

ALTER TABLE "anime_movie_links" DROP CONSTRAINT IF EXISTS "anime_movie_links_link_id_fkey";

ALTER TABLE "anime_serie_links" DROP CONSTRAINT IF EXISTS "anime_serie_links_anime_id_fkey";

ALTER TABLE "anime_serie_links" DROP CONSTRAINT IF EXISTS "anime_serie_links_link_id_fkey";

DROP TABLE IF EXISTS "anime_movie_links";
DROP TABLE IF EXISTS "anime_serie_links";
DROP TABLE IF EXISTS "anime_links";