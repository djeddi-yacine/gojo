CREATE TABLE "anime_episode_servers" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "episode_id" bigserial NOT NULL,
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


CREATE UNIQUE INDEX ON "anime_episode_servers" ("episode_id");

CREATE INDEX ON "anime_serie_videos" ("id");

CREATE INDEX ON "anime_serie_server_sub_videos" ("server_id");

CREATE INDEX ON "anime_serie_server_sub_videos" ("video_id");

CREATE UNIQUE INDEX ON "anime_serie_server_sub_videos" ("server_id", "video_id");

CREATE INDEX ON "anime_serie_server_dub_videos" ("server_id");

CREATE INDEX ON "anime_serie_server_dub_videos" ("video_id");

CREATE UNIQUE INDEX ON "anime_serie_server_dub_videos" ("server_id", "video_id");


ALTER TABLE "anime_episode_servers" ADD FOREIGN KEY ("episode_id") REFERENCES "anime_serie_episodes" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_videos" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_sub_videos" ADD FOREIGN KEY ("server_id") REFERENCES "anime_episode_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_dub_videos" ADD FOREIGN KEY ("server_id") REFERENCES "anime_episode_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_sub_videos" ADD FOREIGN KEY ("video_id") REFERENCES "anime_serie_videos" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_dub_videos" ADD FOREIGN KEY ("video_id") REFERENCES "anime_serie_videos" ("id") ON DELETE CASCADE;