CREATE TABLE "anime_season_official_titles" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "title_text" varchar(150) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_season_short_titles" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "title_text" varchar(150) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_season_translation_titles" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "title_text" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "anime_season_other_titles" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "season_id" bigserial NOT NULL,
  "title_text" varchar(150) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);


CREATE INDEX ON "anime_season_official_titles" ("id", "title_text", "season_id");

CREATE UNIQUE INDEX ON "anime_season_official_titles" ("title_text", "season_id");

CREATE INDEX ON "anime_season_short_titles" ("id", "title_text", "season_id");

CREATE UNIQUE INDEX ON "anime_season_short_titles" ("title_text", "season_id");

CREATE INDEX ON "anime_season_translation_titles" ("id", "title_text", "season_id");

CREATE UNIQUE INDEX ON "anime_season_translation_titles" ("title_text", "season_id");

CREATE INDEX ON "anime_season_other_titles" ("id", "title_text", "season_id");

CREATE UNIQUE INDEX ON "anime_season_other_titles" ("title_text", "season_id");


CREATE INDEX ON "anime_season_official_titles" USING GIN (to_tsvector('pg_catalog.english', lower(translate(title_text, '[:punct:]', ''))));

CREATE INDEX ON "anime_season_short_titles" USING GIN (to_tsvector('pg_catalog.english', lower(translate(title_text, '[:punct:]', ''))));

CREATE INDEX ON "anime_season_translation_titles" USING GIN (to_tsvector('pg_catalog.english', lower(translate(title_text, '[:punct:]', ''))));

CREATE INDEX ON "anime_season_other_titles" USING GIN (to_tsvector('pg_catalog.simple', lower(translate(title_text, '[:punct:]', ''))));



ALTER TABLE "anime_season_official_titles" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_short_titles" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_translation_titles" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;

ALTER TABLE "anime_season_other_titles" ADD FOREIGN KEY ("season_id") REFERENCES "anime_serie_seasons" ("id") ON DELETE CASCADE;
