CREATE TABLE "anime_movie_torrents" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "file_name" varchar NOT NULL,
  "language_id" integer NOT NULL,
  "torrent_hash" varchar NOT NULL,
  "torrent_file" varchar NOT NULL,
  "seeds" integer NOT NULL,
  "peers" integer NOT NULL,
  "leechers" integer NOT NULL,
  "quality" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_movie_server_torrents" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "server_id" bigserial NOT NULL,
  "torrent_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_torrents" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "file_name" varchar NOT NULL,
  "language_id" integer NOT NULL,
  "torrent_hash" varchar NOT NULL,
  "torrent_file" varchar NOT NULL,
  "seeds" integer NOT NULL,
  "peers" integer NOT NULL,
  "leechers" integer NOT NULL,
  "quality" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_serie_server_torrents" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "server_id" bigserial NOT NULL,
  "torrent_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "anime_movie_torrents" ("id");

CREATE INDEX ON "anime_movie_server_torrents" ("server_id");

CREATE INDEX ON "anime_movie_server_torrents" ("torrent_id");

CREATE UNIQUE INDEX ON "anime_movie_server_torrents" ("server_id", "torrent_id");

CREATE INDEX ON "anime_serie_torrents" ("id");

CREATE INDEX ON "anime_serie_server_torrents" ("server_id");

CREATE INDEX ON "anime_serie_server_torrents" ("torrent_id");

CREATE UNIQUE INDEX ON "anime_serie_server_torrents" ("server_id", "torrent_id");


ALTER TABLE "anime_movie_torrents" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_torrents" ADD FOREIGN KEY ("language_id") REFERENCES "languages" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_server_torrents" ADD FOREIGN KEY ("torrent_id") REFERENCES "anime_movie_torrents" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_torrents" ADD FOREIGN KEY ("torrent_id") REFERENCES "anime_serie_torrents" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_movie_server_torrents" ADD FOREIGN KEY ("server_id") REFERENCES "anime_movie_servers" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_serie_server_torrents" ADD FOREIGN KEY ("server_id") REFERENCES "anime_serie_servers" ("id") ON DELETE CASCADE;