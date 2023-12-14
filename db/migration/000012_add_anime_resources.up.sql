CREATE TABLE "anime_resources" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "tvdb_id" integer NOT NULL,
  "tmdb_id" integer NOT NULL,
  "imdb_id" varchar NOT NULL,
  "livechart_id" integer NOT NULL,
  "anime_planet_id" varchar NOT NULL,
  "anisearch_id" integer NOT NULL,
  "anidb_id" integer NOT NULL,
  "kitsu_id" integer NOT NULL,
  "mal_id" integer NOT NULL,
  "notify_moe_id" varchar NOT NULL,
  "anilist_id" integer NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_resources" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "resource_id" bigserial NOT NULL
);

CREATE TABLE "anime_season_resources" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "resource_id" bigserial NOT NULL
);


CREATE INDEX ON "anime_resources" ("id");

CREATE UNIQUE INDEX ON "anime_resources" ("tmdb_id", "imdb_id", "tvdb_id");

CREATE UNIQUE INDEX ON "anime_movie_resources" ("anime_id");

CREATE UNIQUE INDEX ON "anime_movie_resources" ("resource_id");

CREATE UNIQUE INDEX ON "anime_season_resources" ("season_id");

CREATE UNIQUE INDEX ON "anime_season_resources" ("resource_id");


ALTER TABLE "anime_movie_resources" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_resources" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_resources" ADD FOREIGN KEY ("resource_id") REFERENCES "anime_resources" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_resources" ADD FOREIGN KEY ("resource_id") REFERENCES "anime_resources" ("id") ON DELETE CASCADE;
