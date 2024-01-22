ALTER TABLE "anime_episode_metas" DROP CONSTRAINT IF EXISTS "anime_episode_metas_meta_id_fkey";

ALTER TABLE "anime_episode_metas" DROP CONSTRAINT IF EXISTS "anime_episode_metas_episode_id_fkey";

ALTER TABLE "anime_episode_metas" DROP CONSTRAINT IF EXISTS "anime_episode_metas_language_id_fkey";

ALTER TABLE "anime_season_episodes" DROP CONSTRAINT IF EXISTS "anime_season_episodes_season_id_fkey";

ALTER TABLE "anime_season_episodes" DROP CONSTRAINT IF EXISTS "anime_season_episodes_episode_id_fkey";

ALTER TABLE "anime_episodes" DROP CONSTRAINT IF EXISTS "anime_episodes_season_id_fkey";


DROP TABLE IF EXISTS "anime_episode_metas";
DROP TABLE IF EXISTS "anime_season_episodes";
DROP TABLE IF EXISTS "anime_episodes";