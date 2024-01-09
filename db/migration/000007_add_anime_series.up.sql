CREATE TABLE "anime_series" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "original_title" varchar NOT NULL,
  "unique_id" uuid NOT NULL,
  "first_year" integer NOT NULL,
  "last_year" integer NOT NULL,
  "mal_id" integer NOT NULL,
  "tvdb_id" integer NOT NULL,
  "tmdb_id" integer NOT NULL,
  "portrait_poster" varchar NOT NULL,
  "portrait_blur_hash" varchar NOT NULL,
  "landscape_poster" varchar NOT NULL,
  "landscape_blur_hash" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "language_id" integer NOT NULL,
  "meta_id" bigserial NOT NULL
);


CREATE INDEX ON "anime_series" ("id");

CREATE INDEX ON "anime_series" ("original_title");

CREATE INDEX ON "anime_series" ("first_year");

CREATE UNIQUE INDEX ON "anime_series" ("unique_id");

CREATE UNIQUE INDEX ON "anime_series" ("mal_id");

CREATE INDEX ON "anime_serie_metas" ("id");

CREATE UNIQUE INDEX ON "anime_serie_metas" ("anime_id", "language_id");


ALTER TABLE "anime_serie_metas" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_metas" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_metas" ADD FOREIGN KEY ("meta_id") REFERENCES "metas" ("id") ON DELETE CASCADE;
