CREATE TABLE "anime_media" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "media_type" varchar NOT NULL,
  "media_url" varchar NOT NULL,
  "media_author" varchar NOT NULL,
  "media_platform" varchar NOT NULL
);

CREATE TABLE "anime_serie_media" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "media_id" bigserial NOT NULL
);

CREATE TABLE "anime_serie_season_media" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "media_id" bigserial NOT NULL
);

CREATE TABLE "anime_movie_media" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "media_id" bigserial NOT NULL
);



CREATE INDEX ON "anime_media" ("id");

CREATE INDEX ON "anime_serie_media" ("id");

CREATE UNIQUE INDEX ON "anime_serie_media" ("anime_id", "media_id");

CREATE INDEX ON "anime_serie_season_media" ("id");

CREATE UNIQUE INDEX ON "anime_serie_season_media" ("season_id", "media_id");

CREATE INDEX ON "anime_movie_media" ("id");

CREATE UNIQUE INDEX ON "anime_movie_media" ("anime_id", "media_id");



ALTER TABLE "anime_movie_media" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_media" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_media" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_media" ADD FOREIGN KEY ("media_id") REFERENCES "anime_media" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_media" ADD FOREIGN KEY ("media_id") REFERENCES "anime_media" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_media" ADD FOREIGN KEY ("media_id") REFERENCES "anime_media" ("id") ON DELETE CASCADE;
