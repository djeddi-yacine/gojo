CREATE TABLE "anime_characters" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "full_name" varchar NOT NULL,
  "about" varchar NOT NULL,
  "role_playing" varchar NOT NULL,
  "image_url" varchar NOT NULL,
  "image_blur_hash" varchar NOT NULL,
  "pictures" varchar[] NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_character_actors" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "character_id" bigserial NOT NULL,
  "actor_id" bigserial NOT NULL
);

CREATE TABLE "anime_movie_characters" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "character_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_season_characters" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "character_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "anime_characters" ("id");

CREATE UNIQUE INDEX ON "anime_characters" ("full_name", "about");

CREATE UNIQUE INDEX ON "anime_character_actors" ("character_id", "actor_id");

CREATE INDEX ON "anime_movie_characters" ("id");

CREATE UNIQUE INDEX ON "anime_movie_characters" ("anime_id", "character_id");

CREATE INDEX ON "anime_season_characters" ("id");

CREATE UNIQUE INDEX ON "anime_season_characters" ("season_id", "character_id");


ALTER TABLE "anime_movie_characters" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_characters" ADD FOREIGN KEY ("season_id") REFERENCES "anime_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_character_actors" ADD FOREIGN KEY ("actor_id") REFERENCES "actors" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_character_actors" ADD FOREIGN KEY ("character_id") REFERENCES "anime_characters" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_characters" ADD FOREIGN KEY ("character_id") REFERENCES "anime_characters" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_characters" ADD FOREIGN KEY ("character_id") REFERENCES "anime_characters" ("id") ON DELETE CASCADE;
