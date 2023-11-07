ALTER TABLE "anime_serie_season_metas" DROP CONSTRAINT IF EXISTS "anime_serie_season_metas_meta_id_fkey";
ALTER TABLE "anime_serie_season_metas" DROP CONSTRAINT IF EXISTS "anime_serie_season_metas_season_id_fkey";
ALTER TABLE "anime_serie_season_metas" DROP CONSTRAINT IF EXISTS "anime_serie_season_metas_language_id_fkey";

ALTER TABLE "anime_serie_season_episodes" DROP CONSTRAINT IF EXISTS "anime_serie_season_episodes_season_id_fkey";
ALTER TABLE "anime_serie_season_episodes" DROP CONSTRAINT IF EXISTS "anime_serie_season_episodes_episode_id_fkey";

ALTER TABLE "anime_serie_seasons" DROP CONSTRAINT IF EXISTS "anime_serie_seasons_anime_id_fkey";

ALTER TABLE "anime_serie_episode_metas" DROP CONSTRAINT IF EXISTS "anime_serie_episode_metas_meta_id_fkey";
ALTER TABLE "anime_serie_episode_metas" DROP CONSTRAINT IF EXISTS "anime_serie_episode_metas_episode_id_fkey";
ALTER TABLE "anime_serie_episode_metas" DROP CONSTRAINT IF EXISTS "anime_serie_episode_metas_language_id_fkey";

ALTER TABLE "anime_serie_episode_servers" DROP CONSTRAINT IF EXISTS "anime_serie_episode_servers_server_id_fkey";
ALTER TABLE "anime_serie_episode_servers" DROP CONSTRAINT IF EXISTS "anime_serie_episode_servers_episode_id_fkey";

ALTER TABLE "anime_serie_episodes" DROP CONSTRAINT IF EXISTS "anime_serie_episodes_season_id_fkey";

ALTER TABLE "anime_serie_server_sub_videos" DROP CONSTRAINT IF EXISTS "anime_serie_server_sub_videos_server_id_fkey";
ALTER TABLE "anime_serie_server_sub_videos" DROP CONSTRAINT IF EXISTS "anime_serie_server_sub_videos_video_id_fkey";

ALTER TABLE "anime_serie_server_dub_videos" DROP CONSTRAINT IF EXISTS "anime_serie_server_dub_videos_server_id_fkey";
ALTER TABLE "anime_serie_server_dub_videos" DROP CONSTRAINT IF EXISTS "anime_serie_server_dub_videos_video_id_fkey";

ALTER TABLE "anime_serie_servers" DROP CONSTRAINT IF EXISTS "anime_serie_servers_episode_id_fkey";

ALTER TABLE "anime_serie_videos" DROP CONSTRAINT IF EXISTS "anime_serie_videos_language_id_fkey";

DROP TABLE IF EXISTS "anime_serie_season_metas";
DROP TABLE IF EXISTS "anime_serie_season_episodes";
DROP TABLE IF EXISTS "anime_serie_seasons";
DROP TABLE IF EXISTS "anime_serie_episode_metas";
DROP TABLE IF EXISTS "anime_serie_episode_servers";
DROP TABLE IF EXISTS "anime_serie_episodes";
DROP TABLE IF EXISTS "anime_serie_server_sub_videos";
DROP TABLE IF EXISTS "anime_serie_server_dub_videos";
DROP TABLE IF EXISTS "anime_serie_servers";
DROP TABLE IF EXISTS "anime_serie_videos";