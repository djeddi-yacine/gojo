ALTER TABLE "anime_movie_torrents" DROP CONSTRAINT IF EXISTS "anime_movie_torrents_language_id_fkey";

ALTER TABLE "anime_serie_torrents" DROP CONSTRAINT IF EXISTS "anime_serie_torrents_language_id_fkey";

ALTER TABLE "anime_movie_server_torrents" DROP CONSTRAINT IF EXISTS "anime_movie_server_torrents_torrent_id_fkey";

ALTER TABLE "anime_movie_server_torrents" DROP CONSTRAINT IF EXISTS "anime_movie_server_torrents_server_id_fkey";

ALTER TABLE "anime_serie_server_torrents" DROP CONSTRAINT IF EXISTS "anime_serie_server_torrents_torrent_id_fkey";

ALTER TABLE "anime_serie_server_torrents" DROP CONSTRAINT IF EXISTS "anime_serie_server_torrents_server_id_fkey";

DROP TABLE IF EXISTS "anime_movie_server_torrents";
DROP TABLE IF EXISTS "anime_movie_torrents";
DROP TABLE IF EXISTS "anime_serie_server_torrents";
DROP TABLE IF EXISTS "anime_serie_torrents";