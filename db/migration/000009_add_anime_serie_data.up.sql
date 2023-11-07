CREATE TABLE "anime_serie_seasons" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "season_number" integer UNIQUE NOT NULL,
  "portriat_poster" varchar NOT NULL,
  "portriat_blur_hash" varchar NOT NULL,
  "landscape_poster" varchar NOT NULL,
  "landscape_blur_hash" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_season_metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "meta_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_episodes" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "episode_number" integer UNIQUE NOT NULL,
  "season_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_episode_metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "episode_id" bigserial NOT NULL,
  "meta_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_season_episodes" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "episode_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_servers" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "episode_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_episode_servers" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "episode_id" bigserial NOT NULL,
  "server_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_videos" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "language_id" integer NOT NULL,
  "autority" varchar NOT NULL,
  "referer" varchar NOT NULL,
  "link" varchar NOT NULL,
  "quality" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_server_sub_videos" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "server_id" bigserial NOT NULL,
  "video_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_server_dub_videos" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "server_id" bigserial NOT NULL,
  "video_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);



CREATE INDEX ON "anime_serie_seasons" ("id");

CREATE UNIQUE INDEX ON "anime_serie_seasons" ("anime_id", "season_number");

CREATE INDEX ON "anime_serie_season_metas" ("id");

CREATE UNIQUE INDEX ON "anime_serie_season_metas" ("season_id", "meta_id");

CREATE INDEX ON "anime_serie_episodes" ("id");

CREATE UNIQUE INDEX ON "anime_serie_episodes" ("episode_number", "season_id");

CREATE INDEX ON "anime_serie_episode_metas" ("id");

CREATE UNIQUE INDEX ON "anime_serie_episode_metas" ("episode_id", "meta_id");

CREATE INDEX ON "anime_serie_season_episodes" ("season_id");

CREATE INDEX ON "anime_serie_season_episodes" ("episode_id");

CREATE UNIQUE INDEX ON "anime_serie_season_episodes" ("season_id", "episode_id");

CREATE INDEX ON "anime_serie_servers" ("id");

CREATE INDEX ON "anime_serie_episode_servers" ("server_id");

CREATE INDEX ON "anime_serie_episode_servers" ("episode_id");

CREATE UNIQUE INDEX ON "anime_serie_episode_servers" ("episode_id", "server_id");

CREATE INDEX ON "anime_serie_server_sub_videos" ("server_id");

CREATE INDEX ON "anime_serie_server_sub_videos" ("video_id");

CREATE UNIQUE INDEX ON "anime_serie_server_sub_videos" ("server_id", "video_id");

CREATE INDEX ON "anime_serie_server_dub_videos" ("server_id");

CREATE INDEX ON "anime_serie_server_dub_videos" ("video_id");

CREATE UNIQUE INDEX ON "anime_serie_server_dub_videos" ("server_id", "video_id");

CREATE INDEX ON "anime_serie_videos" ("id");



ALTER TABLE "anime_serie_seasons" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_videos" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_servers" ADD FOREIGN KEY ("episode_id") REFERENCES "anime_serie_episodes" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_episodes" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_metas" ADD FOREIGN KEY ("meta_id") REFERENCES "metas" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_episode_metas" ADD FOREIGN KEY ("meta_id") REFERENCES "metas" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_metas" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_episode_metas" ADD FOREIGN KEY ("episode_id") REFERENCES "anime_serie_episodes" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_episodes" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_episodes" ADD FOREIGN KEY ("episode_id") REFERENCES "anime_serie_episodes" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_episode_servers" ADD FOREIGN KEY ("episode_id") REFERENCES "anime_serie_episodes" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_episode_servers" ADD FOREIGN KEY ("server_id") REFERENCES "anime_serie_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_sub_videos" ADD FOREIGN KEY ("server_id") REFERENCES "anime_serie_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_dub_videos" ADD FOREIGN KEY ("server_id") REFERENCES "anime_serie_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_sub_videos" ADD FOREIGN KEY ("video_id") REFERENCES "anime_serie_videos" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_dub_videos" ADD FOREIGN KEY ("video_id") REFERENCES "anime_serie_videos" ("id") ON DELETE CASCADE;