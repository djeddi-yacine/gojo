CREATE TABLE "anime_tags" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "tag" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_tags" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "tag_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_season_tags" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "tag_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "anime_tags" ("id");

CREATE UNIQUE INDEX ON "anime_tags" ("tag");

CREATE INDEX ON "anime_movie_tags" ("id");

CREATE UNIQUE INDEX ON "anime_movie_tags" ("anime_id", "tag_id");

CREATE INDEX ON "anime_season_tags" ("id");

CREATE UNIQUE INDEX ON "anime_season_tags" ("season_id", "tag_id");


ALTER TABLE "anime_movie_tags" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_tags" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "anime_tags" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_tags" ADD FOREIGN KEY ("tag_id") REFERENCES "anime_tags" ("id") ON DELETE CASCADE;