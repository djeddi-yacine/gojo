CREATE TABLE "anime_serie_seasons" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "season_original_title" varchar NOT NULL,
  "release_year" integer NOT NULL,
  "aired" timestamptz NOT NULL,
  "portrait_poster" varchar NOT NULL,
  "portrait_blur_hash" varchar NOT NULL,
  "rating" varchar NOT NULL DEFAULT ('PG-13 - Teens 13 or older'),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_season_metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "language_id" integer NOT NULL,
  "meta_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_season_studios" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "studio_id" integer NOT NULL
);

CREATE TABLE "anime_season_genres" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "genre_id" integer NOT NULL
);


CREATE INDEX ON "anime_serie_seasons" ("id");

CREATE INDEX ON "anime_serie_seasons" ("release_year");

CREATE UNIQUE INDEX ON "anime_serie_seasons" ("anime_id", "season_original_title", "release_year");

CREATE INDEX ON "anime_season_metas" ("id");

CREATE UNIQUE INDEX ON "anime_season_metas" ("season_id", "language_id");

CREATE INDEX ON "anime_season_studios" ("id");

CREATE UNIQUE INDEX ON "anime_season_studios" ("season_id", "studio_id");

CREATE INDEX ON "anime_season_genres" ("id");

CREATE UNIQUE INDEX ON "anime_season_genres" ("season_id", "genre_id");


ALTER TABLE "anime_serie_seasons" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_metas" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_metas" ADD FOREIGN KEY ("meta_id") REFERENCES "metas" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_metas" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_genres" ADD FOREIGN KEY ("genre_id") REFERENCES "genres" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_studios" ADD FOREIGN KEY ("studio_id") REFERENCES "studios" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_studios" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_genres" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;
