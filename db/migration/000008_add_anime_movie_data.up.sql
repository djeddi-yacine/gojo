CREATE TABLE "anime_movie_servers" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "anime_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_server_sub_videos" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "server_id" bigserial NOT NULL,
  "video_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_server_dub_videos" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "server_id" bigserial NOT NULL,
  "video_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_videos" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "language_id" integer NOT NULL,
  "authority" varchar NOT NULL,
  "referer" varchar NOT NULL,
  "link" varchar NOT NULL,
  "quality" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);



CREATE INDEX ON "anime_movie_servers" ("id");

CREATE UNIQUE INDEX ON "anime_movie_server_sub_videos" ("server_id", "video_id");

CREATE UNIQUE INDEX ON "anime_movie_server_dub_videos" ("server_id", "video_id");

CREATE INDEX ON "anime_movie_videos" ("id");



ALTER TABLE "anime_movie_servers" ADD FOREIGN KEY ("anime_id") REFERENCES "anime_movies" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_videos" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_server_sub_videos" ADD FOREIGN KEY ("server_id") REFERENCES "anime_movie_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_server_dub_videos" ADD FOREIGN KEY ("server_id") REFERENCES "anime_movie_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_server_sub_videos" ADD FOREIGN KEY ("video_id") REFERENCES "anime_movie_videos" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_server_dub_videos" ADD FOREIGN KEY ("video_id") REFERENCES "anime_movie_videos" ("id") ON DELETE CASCADE;