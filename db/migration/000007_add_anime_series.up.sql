CREATE TABLE "anime_series" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "original_title" varchar NOT NULL,
  "aired" timestamptz NOT NULL,
  "release_year" integer NOT NULL,
  "duration" interval NOT NULL DEFAULT ('00h 00m 00s'),
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_studios" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "studio_id" integer
);

CREATE TABLE "anime_serie_genres" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "genre_id" integer
);

CREATE TABLE "anime_serie_metas" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "language_id" integer NOT NULL,
  "meta_id" bigserial NOT NULL
);


CREATE INDEX ON "anime_series" ("id");

CREATE INDEX ON "anime_series" ("original_title");

CREATE INDEX ON "anime_series" ("release_year");

CREATE UNIQUE INDEX ON "anime_series" ("original_title", "duration", "aired");

CREATE INDEX ON "anime_serie_studios" ("anime_id");

CREATE INDEX ON "anime_serie_studios" ("studio_id");

CREATE INDEX ON "anime_serie_genres" ("anime_id");

CREATE INDEX ON "anime_serie_genres" ("genre_id");

CREATE INDEX ON "anime_serie_metas" ("id");

CREATE UNIQUE INDEX ON "anime_serie_metas" ("anime_id", "language_id");


ALTER TABLE "anime_serie_studios" ADD FOREIGN KEY ("studio_id") REFERENCES "studios" ("id");

ALTER TABLE "anime_serie_genres" ADD FOREIGN KEY ("genre_id") REFERENCES "genres" ("id");

ALTER TABLE "anime_serie_metas" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_studios" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_genres" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_metas" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_metas" ADD FOREIGN KEY ("meta_id") REFERENCES "metas" ("id") ON DELETE CASCADE;
