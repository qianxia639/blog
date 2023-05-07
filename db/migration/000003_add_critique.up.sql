CREATE TABLE "critiques" (
  "id" bigserial PRIMARY KEY,
  "owner_id" int8 NOT NULL,
  "parent_id" int8 NOT NULL DEFAULT 0,
  "nickname" varchar(20) NOT NULL,
  "avatar" varchar(255) NOT NULL DEFAULT '',
  "content" varchar(255) NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

COMMENT ON COLUMN "critiques"."id" IS '主键';

COMMENT ON COLUMN "critiques"."owner_id" IS '博客Id';

COMMENT ON COLUMN "critiques"."parent_id" IS '父评论Id';

COMMENT ON COLUMN "critiques"."nickname" IS '昵称';

COMMENT ON COLUMN "critiques"."avatar" IS '头像';

COMMENT ON COLUMN "critiques"."content" IS '评论内容';

COMMENT ON COLUMN "critiques"."created_at" IS '创建时间';

ALTER TABLE "critiques" ADD FOREIGN KEY ("owner_id") REFERENCES "critiques" ("id");