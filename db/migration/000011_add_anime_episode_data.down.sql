ALTER TABLE "anime_episode_server_sub_videos" DROP CONSTRAINT IF EXISTS "anime_episode_server_sub_videos_server_id_fkey";

ALTER TABLE "anime_episode_server_sub_videos" DROP CONSTRAINT IF EXISTS "anime_episode_server_sub_videos_video_id_fkey";

ALTER TABLE "anime_episode_server_dub_videos" DROP CONSTRAINT IF EXISTS "anime_episode_server_dub_videos_server_id_fkey";

ALTER TABLE "anime_episode_server_dub_videos" DROP CONSTRAINT IF EXISTS "anime_episode_server_dub_videos_video_id_fkey";

ALTER TABLE "anime_episode_servers" DROP CONSTRAINT IF EXISTS "anime_episode_servers_episode_id_fkey";

ALTER TABLE "anime_episode_videos" DROP CONSTRAINT IF EXISTS "anime_episode_videos_language_id_fkey";


DROP TABLE IF EXISTS "anime_episode_server_sub_videos";
DROP TABLE IF EXISTS "anime_episode_server_dub_videos";
DROP TABLE IF EXISTS "anime_episode_videos";
DROP TABLE IF EXISTS "anime_episode_servers";