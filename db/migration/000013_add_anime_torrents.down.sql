ALTER TABLE "anime_movie_torrents" DROP CONSTRAINT IF EXISTS "anime_movie_torrents_language_id_fkey";

ALTER TABLE "anime_episode_torrents" DROP CONSTRAINT IF EXISTS "anime_episode_torrents_language_id_fkey";

ALTER TABLE "anime_movie_server_sub_torrents" DROP CONSTRAINT IF EXISTS "anime_movie_server_sub_torrents_torrent_id_fkey";

ALTER TABLE "anime_movie_server_sub_torrents" DROP CONSTRAINT IF EXISTS "anime_movie_server_sub_torrents_server_id_fkey";

ALTER TABLE "anime_movie_server_dub_torrents" DROP CONSTRAINT IF EXISTS "anime_movie_server_dub_torrents_torrent_id_fkey";

ALTER TABLE "anime_movie_server_dub_torrents" DROP CONSTRAINT IF EXISTS "anime_movie_server_dub_torrents_server_id_fkey";

ALTER TABLE "anime_episode_server_sub_torrents" DROP CONSTRAINT IF EXISTS "anime_episode_server_sub_torrents_torrent_id_fkey";

ALTER TABLE "anime_episode_server_sub_torrents" DROP CONSTRAINT IF EXISTS "anime_episode_server_sub_torrents_server_id_fkey";

ALTER TABLE "anime_episode_server_dub_torrents" DROP CONSTRAINT IF EXISTS "anime_episode_server_dub_torrents_torrent_id_fkey";

ALTER TABLE "anime_episode_server_dub_torrents" DROP CONSTRAINT IF EXISTS "anime_episode_server_dub_torrents_server_id_fkey";

DROP TABLE IF EXISTS "anime_movie_server_dub_torrents";
DROP TABLE IF EXISTS "anime_movie_server_sub_torrents";
DROP TABLE IF EXISTS "anime_movie_torrents";
DROP TABLE IF EXISTS "anime_episode_server_dub_torrents";
DROP TABLE IF EXISTS "anime_episode_server_sub_torrents";
DROP TABLE IF EXISTS "anime_episode_torrents";