CREATE TABLE "anime_images" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "image_type" integer NOT NULL,
  "image_host" varchar NOT NULL,
  "image_url" varchar NOT NULL,
  "image_thumbnails" varchar NOT NULL,
  "image_blurhash" varchar NOT NULL,
  "image_height" integer NOT NULL,
  "image_width" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_images" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "image_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_images" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "image_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_season_images" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "image_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);



CREATE INDEX ON "anime_images" ("id");

CREATE INDEX ON "anime_serie_images" ("id");

CREATE UNIQUE INDEX ON "anime_serie_images" ("anime_id", "image_id");

CREATE INDEX ON "anime_serie_season_images" ("id");

CREATE UNIQUE INDEX ON "anime_serie_season_images" ("season_id", "image_id");

CREATE INDEX ON "anime_movie_images" ("id");

CREATE UNIQUE INDEX ON "anime_movie_images" ("anime_id", "image_id");



ALTER TABLE "anime_movie_images" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_images" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_images" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_images" ADD FOREIGN KEY ("image_id") REFERENCES "anime_images" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_images" ADD FOREIGN KEY ("image_id") REFERENCES "anime_images" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_season_images" ADD FOREIGN KEY ("image_id") REFERENCES "anime_images" ("id") ON DELETE CASCADE;
