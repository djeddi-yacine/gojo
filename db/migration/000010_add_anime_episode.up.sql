CREATE TABLE "anime_serie_episodes" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "episode_number" integer NOT NULL,
  "episode_original_title" varchar NOT NULL,
  "aired" timestamptz NOT NULL,
  "rating" varchar NOT NULL DEFAULT ('PG-13 - Teens 13 or older'),
  "duration" interval NOT NULL DEFAULT ('00h 00m 00s'),
  "thumbnails" varchar NOT NULL,
  "thumbnails_blur_hash" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_episode_metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "episode_id" bigserial NOT NULL,
  "language_id" integer NOT NULL,
  "meta_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_season_episodes" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "episode_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "anime_serie_episodes" ("id");

CREATE UNIQUE INDEX ON "anime_serie_episodes" ("episode_number", "season_id");

CREATE INDEX ON "anime_episode_metas" ("id");

CREATE UNIQUE INDEX ON "anime_episode_metas" ("episode_id", "language_id");

CREATE UNIQUE INDEX ON "anime_season_episodes" ("season_id", "episode_id");


ALTER TABLE "anime_serie_episodes" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_episode_metas" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_episode_metas" ADD FOREIGN KEY ("meta_id") REFERENCES "metas" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_episode_metas" ADD FOREIGN KEY ("episode_id") REFERENCES "anime_serie_episodes" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_episodes" ADD FOREIGN KEY ("episode_id") REFERENCES "anime_serie_episodes" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_episodes" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;