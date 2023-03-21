CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "blog_id" int8 NOT NULL,
  "comment_id" int8 NOT NULL,
  "nickname" varchar(20) NOT NULL,
  "avatar" varchar(255) NOT NULL DEFAULT '',
  "content" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

COMMENT ON COLUMN "comments"."id" IS '主键';

COMMENT ON COLUMN "comments"."blog_id" IS '博客Id';

COMMENT ON COLUMN "comments"."comment_id" IS '父评论Id';

COMMENT ON COLUMN "comments"."nickname" IS '昵称';

COMMENT ON COLUMN "comments"."avatar" IS '头像';

COMMENT ON COLUMN "comments"."content" IS '评论内容';

COMMENT ON COLUMN "comments"."created_at" IS '创建时间';

ALTER TABLE "comments" ADD FOREIGN KEY ("blog_id") REFERENCES "blogs" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("comment_id") REFERENCES "comments" ("id");