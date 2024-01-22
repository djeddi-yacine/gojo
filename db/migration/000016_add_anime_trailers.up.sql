CREATE TABLE "anime_trailers" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "is_official" bool NOT NULL DEFAULT false,
  "host_name" varchar NOT NULL,
  "host_key" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_trailers" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "trailer_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_trailers" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "trailer_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_season_trailers" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "trailer_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "anime_trailers" ("id");

CREATE INDEX ON "anime_movie_trailers" ("id");

CREATE UNIQUE INDEX ON "anime_movie_trailers" ("anime_id", "trailer_id");

CREATE INDEX ON "anime_serie_trailers" ("id");

CREATE UNIQUE INDEX ON "anime_serie_trailers" ("anime_id", "trailer_id");

CREATE INDEX ON "anime_season_trailers" ("id");

CREATE UNIQUE INDEX ON "anime_season_trailers" ("season_id", "trailer_id");


ALTER TABLE "anime_movie_trailers" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_trailers" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_series" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_trailers" ADD FOREIGN KEY ("season_id") REFERENCES "anime_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_trailers" ADD FOREIGN KEY ("trailer_id") REFERENCES "anime_trailers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_trailers" ADD FOREIGN KEY ("trailer_id") REFERENCES "anime_trailers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_trailers" ADD FOREIGN KEY ("trailer_id") REFERENCES "anime_trailers" ("id") ON DELETE CASCADE;
