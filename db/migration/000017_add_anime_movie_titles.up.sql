CREATE EXTENSION IF NOT EXISTS pg_trgm;

CREATE TABLE "anime_movie_official_titles" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "title_text" varchar(150) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_short_titles" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "title_text" varchar(150) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_other_titles" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "title_text" varchar(150) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_translation_titles" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "title_text" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "anime_movie_official_titles" ("id", "title_text", "anime_id");

CREATE INDEX ON "anime_movie_short_titles" ("id", "title_text", "anime_id");

CREATE INDEX ON "anime_movie_translation_titles" ("id", "title_text", "anime_id");

CREATE INDEX ON "anime_movie_other_titles" ("id", "title_text", "anime_id");


CREATE INDEX ON "anime_movie_official_titles" USING GIN (to_tsvector('pg_catalog.english', lower(translate(title_text, '[:punct:]', ''))));

CREATE INDEX ON "anime_movie_short_titles" USING GIN (to_tsvector('pg_catalog.english', lower(translate(title_text, '[:punct:]', ''))));

CREATE INDEX ON "anime_movie_other_titles" USING GIN (to_tsvector('pg_catalog.english', lower(translate(title_text, '[:punct:]', ''))));

CREATE INDEX ON "anime_movie_translation_titles" USING GIN (to_tsvector('pg_catalog.simple', lower(translate(title_text, '[:punct:]', ''))));


ALTER TABLE "anime_movie_official_titles" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_short_titles" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_translation_titles" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_other_titles" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;