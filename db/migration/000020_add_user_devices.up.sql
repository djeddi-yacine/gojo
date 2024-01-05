CREATE TABLE "devices" (
  "id" uuid UNIQUE PRIMARY KEY NOT NULL,
  "operating_system" varchar NOT NULL,
  "mac_address" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "is_banned" boolean NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "user_devices" (
  "id" bigserial UNIQUE PRIMARY KEY NOT NULL,
  "user_id" bigserial NOT NULL,
  "device_id" uuid NOT NULL
);


CREATE INDEX ON "devices" ("id");

CREATE UNIQUE INDEX ON "devices" ("operating_system", "mac_address", "client_ip");

CREATE UNIQUE INDEX ON "user_devices" ("user_id", "device_id");


ALTER TABLE "user_devices" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "user_devices" ADD FOREIGN KEY ("device_id") REFERENCES "devices" ("id") ON DELETE CASCADE;