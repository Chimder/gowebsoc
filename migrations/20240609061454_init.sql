-- +goose Up
-- create "channels" table
CREATE TABLE "public"."channels" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_channels_name" UNIQUE ("name")
);
-- create index "idx_channels_deleted_at" to table: "channels"
CREATE INDEX "idx_channels_deleted_at" ON "public"."channels" ("deleted_at");
-- create "podchannels" table
CREATE TABLE "public"."podchannels" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NOT NULL,
  "type" text NULL DEFAULT 'text',
  "channel_id" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_channels_pod_channels" FOREIGN KEY ("channel_id") REFERENCES "public"."channels" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_podchannels_deleted_at" to table: "podchannels"
CREATE INDEX "idx_podchannels_deleted_at" ON "public"."podchannels" ("deleted_at");
-- create "users" table
CREATE TABLE "public"."users" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "name" text NOT NULL,
  "email" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "uni_users_email" UNIQUE ("email"),
  CONSTRAINT "uni_users_name" UNIQUE ("name")
);
-- create index "idx_users_deleted_at" to table: "users"
CREATE INDEX "idx_users_deleted_at" ON "public"."users" ("deleted_at");
-- create "messages" table
CREATE TABLE "public"."messages" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "content" text NOT NULL,
  "author_id" bigint NOT NULL,
  "podchannel_id" bigint NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_podchannels_messages" FOREIGN KEY ("podchannel_id") REFERENCES "public"."podchannels" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION,
  CONSTRAINT "fk_users_messages" FOREIGN KEY ("author_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_messages_deleted_at" to table: "messages"
CREATE INDEX "idx_messages_deleted_at" ON "public"."messages" ("deleted_at");

-- +goose Down
-- reverse: create index "idx_messages_deleted_at" to table: "messages"
DROP INDEX "public"."idx_messages_deleted_at";
-- reverse: create "messages" table
DROP TABLE "public"."messages";
-- reverse: create index "idx_users_deleted_at" to table: "users"
DROP INDEX "public"."idx_users_deleted_at";
-- reverse: create "users" table
DROP TABLE "public"."users";
-- reverse: create index "idx_podchannels_deleted_at" to table: "podchannels"
DROP INDEX "public"."idx_podchannels_deleted_at";
-- reverse: create "podchannels" table
DROP TABLE "public"."podchannels";
-- reverse: create index "idx_channels_deleted_at" to table: "channels"
DROP INDEX "public"."idx_channels_deleted_at";
-- reverse: create "channels" table
DROP TABLE "public"."channels";
