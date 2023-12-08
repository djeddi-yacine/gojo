CREATE TABLE "anime_links" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "official_website" varchar NOT NULL,
  "wikipedia_url" varchar NOT NULL,
  "crunchyroll_url" varchar NOT NULL,
  "social_media" varchar[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_links" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial UNIQUE NOT NULL,
  "link_id" bigserial UNIQUE NOT NULL
);

CREATE TABLE "anime_movie_links" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial UNIQUE NOT NULL,
  "link_id" bigserial UNIQUE NOT NULL
);


CREATE INDEX ON "anime_links" ("id");

CREATE UNIQUE INDEX ON "anime_links" ("wikipedia_url", "official_website");

CREATE INDEX ON "anime_serie_links" ("anime_id");

CREATE INDEX ON "anime_serie_links" ("link_id");

CREATE INDEX ON "anime_movie_links" ("anime_id");

CREATE INDEX ON "anime_movie_links" ("link_id");


ALTER TABLE "anime_movie_links" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_links" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_links" ADD FOREIGN KEY ("link_id") REFERENCES "anime_links" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_links" ADD FOREIGN KEY ("link_id") REFERENCES "anime_links" ("id") ON DELETE CASCADE;
