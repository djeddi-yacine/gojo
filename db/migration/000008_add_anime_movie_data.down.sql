ALTER TABLE "anime_movie_servers" DROP CONSTRAINT IF EXISTS "anime_movie_servers_anime_id_fkey";
ALTER TABLE "anime_movie_servers" DROP CONSTRAINT IF EXISTS "anime_movie_servers_sub_id_fkey";
ALTER TABLE "anime_movie_servers" DROP CONSTRAINT IF EXISTS "anime_movie_servers_dub_id_fkey";

ALTER TABLE "anime_movie_server_sub_videos" DROP CONSTRAINT IF EXISTS "anime_movie_server_sub_videos_server_id_fkey";
ALTER TABLE "anime_movie_server_sub_videos" DROP CONSTRAINT IF EXISTS "anime_movie_server_sub_videos_video_id_fkey";

ALTER TABLE "anime_movie_server_dub_videos" DROP CONSTRAINT IF EXISTS "anime_movie_server_dub_videos_server_id_fkey";
ALTER TABLE "anime_movie_server_dub_videos" DROP CONSTRAINT IF EXISTS "anime_movie_server_dub_videos_video_id_fkey";

ALTER TABLE "anime_movie_videos" DROP CONSTRAINT IF EXISTS "anime_movie_videos_language_id_fkey";

DROP TABLE IF EXISTS "anime_movie_server_sub_videos";
DROP TABLE IF EXISTS "anime_movie_server_dub_videos";
DROP TABLE IF EXISTS "anime_movie_servers";
DROP TABLE IF EXISTS "anime_movie_videos";