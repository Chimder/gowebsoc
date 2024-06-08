-- +goose Up
-- drop "blog_posts" table
DROP TABLE "public"."blog_posts";

-- +goose Down
-- reverse: drop "blog_posts" table
CREATE TABLE "public"."blog_posts" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "user_id" bigint NOT NULL,
  PRIMARY KEY ("id")
);
CREATE INDEX "idx_blog_posts_deleted_at" ON "public"."blog_posts" ("deleted_at");
