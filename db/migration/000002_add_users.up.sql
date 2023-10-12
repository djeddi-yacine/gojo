CREATE TABLE "users" (
  "id" BIGSERIAL UNIQUE NOT NULL,
  "username" varchar UNIQUE NOT NULL,
  "email" varchar NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "password_changed_at" timestamptz NOT NULL DEFAULT (now()),
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("id", "username")
);


CREATE INDEX ON "users" ("username");

CREATE INDEX ON "users" ("full_name");