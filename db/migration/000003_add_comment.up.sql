CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "owner_id" int8 NOT NULL,
  "parent_id" int8 NOT NULL DEFAULT 0,
  "nickname" varchar(20) NOT NULL,
  "avatar" varchar(255) NOT NULL DEFAULT '',
  "content" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

COMMENT ON COLUMN "comments"."id" IS '主键';

COMMENT ON COLUMN "comments"."owner_id" IS '博客Id';

COMMENT ON COLUMN "comments"."parent_id" IS '父评论Id';

COMMENT ON COLUMN "comments"."nickname" IS '昵称';

COMMENT ON COLUMN "comments"."avatar" IS '头像';

COMMENT ON COLUMN "comments"."content" IS '评论内容';

COMMENT ON COLUMN "comments"."created_at" IS '创建时间';

ALTER TABLE "comments" ADD FOREIGN KEY ("owner_id") REFERENCES "blogs" ("id");