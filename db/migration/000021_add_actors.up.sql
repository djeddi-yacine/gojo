CREATE TABLE "actors" (
  "id" BIGSERIAL UNIQUE PRIMARY KEY NOT NULL,
  "full_name" varchar NOT NULL,
  "gender" varchar NOT NULL,
  "biography" varchar NOT NULL,
  "born" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE INDEX ON "actors" ("id");

CREATE UNIQUE INDEX ON "actors" ("full_name", "born");