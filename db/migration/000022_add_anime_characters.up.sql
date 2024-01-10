CREATE TABLE "anime_characters" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "actor_id" bigserial NOT NULL,
  "full_name" varchar NOT NULL,
  "description" varchar NOT NULL,
  "image" varchar NOT NULL,
  "image_blur_hash" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_characters" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "character_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_characters" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "character_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "anime_characters" ("id");

CREATE UNIQUE INDEX ON "anime_characters" ("actor_id", "full_name");

CREATE INDEX ON "anime_movie_characters" ("id");

CREATE UNIQUE INDEX ON "anime_movie_characters" ("anime_id", "character_id");

CREATE INDEX ON "anime_serie_characters" ("id");

CREATE UNIQUE INDEX ON "anime_serie_characters" ("anime_id", "character_id");


ALTER TABLE "anime_movie_characters" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_characters" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_characters" ADD FOREIGN KEY ("actor_id") REFERENCES "actors" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_characters" ADD FOREIGN KEY ("character_id") REFERENCES "anime_characters" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_characters" ADD FOREIGN KEY ("character_id") REFERENCES "anime_characters" ("id") ON DELETE CASCADE;
