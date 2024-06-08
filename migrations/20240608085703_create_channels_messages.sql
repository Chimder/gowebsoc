-- +goose Up
-- modify "blog_posts" table
ALTER TABLE "public"."blog_posts" DROP CONSTRAINT "fk_users_blog_posts";
-- modify "users" table
ALTER TABLE "public"."users" DROP COLUMN "first_name", DROP COLUMN "last_name", ADD COLUMN "name" text NOT NULL, ADD CONSTRAINT "uni_users_name" UNIQUE ("name");
-- create "messages" table
CREATE TABLE "public"."messages" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "content" text NOT NULL,
  "author_id" bigint NOT NULL,
  "channel_id" text NOT NULL,
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_users_messages" FOREIGN KEY ("author_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_messages_deleted_at" to table: "messages"
CREATE INDEX "idx_messages_deleted_at" ON "public"."messages" ("deleted_at");
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
  "channel_id" bigint NOT NULL,
  "name" text NOT NULL,
  "type" text NULL DEFAULT 'text',
  PRIMARY KEY ("id"),
  CONSTRAINT "fk_channels_pod_channels" FOREIGN KEY ("channel_id") REFERENCES "public"."channels" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION
);
-- create index "idx_podchannels_deleted_at" to table: "podchannels"
CREATE INDEX "idx_podchannels_deleted_at" ON "public"."podchannels" ("deleted_at");

-- +goose Down
-- reverse: create index "idx_podchannels_deleted_at" to table: "podchannels"
DROP INDEX "public"."idx_podchannels_deleted_at";
-- reverse: create "podchannels" table
DROP TABLE "public"."podchannels";
-- reverse: create index "idx_channels_deleted_at" to table: "channels"
DROP INDEX "public"."idx_channels_deleted_at";
-- reverse: create "channels" table
DROP TABLE "public"."channels";
-- reverse: create index "idx_messages_deleted_at" to table: "messages"
DROP INDEX "public"."idx_messages_deleted_at";
-- reverse: create "messages" table
DROP TABLE "public"."messages";
-- reverse: modify "users" table
ALTER TABLE "public"."users" DROP CONSTRAINT "uni_users_name", DROP COLUMN "name", ADD COLUMN "last_name" text NOT NULL, ADD COLUMN "first_name" text NOT NULL;
-- reverse: modify "blog_posts" table
ALTER TABLE "public"."blog_posts" ADD
 CONSTRAINT "fk_users_blog_posts" FOREIGN KEY ("user_id") REFERENCES "public"."users" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
