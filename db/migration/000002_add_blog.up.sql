CREATE TABLE "blogs" (
  "id" bigserial PRIMARY KEY,
  "owner_id" int8 NOT NULL,
  "title" varchar(50) NOT NULL UNIQUE,
  "content" text NOT NULL,
  "image" varchar(255) NOT NULL,
  "views" int4 DEFAULT 0 NOT NULL,
  "is_reward" bool NOT NULL DEFAULT false,
  "is_critique" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT '1970-01-01 00:00:00',
  "updated_at" timestamptz NOT NULL DEFAULT '1970-01-01 00:00:00'
);

COMMENT ON COLUMN "blogs"."id" IS '主键';

COMMENT ON COLUMN "blogs"."owner_id" IS '创建者Id';

COMMENT ON COLUMN "blogs"."title" IS '标题';

COMMENT ON COLUMN "blogs"."content" IS '内容';

COMMENT ON COLUMN "blogs"."image" IS '图片链接';

COMMENT ON COLUMN "blogs"."views" IS '浏览次数';

COMMENT ON COLUMN "blogs"."is_reward" IS '是否开启打赏(t: 是,f: 否)';

COMMENT ON COLUMN "blogs"."is_critique" IS '是否开启评论(t: 是,f: 否)';

COMMENT ON COLUMN "blogs"."created_at" IS '创建时间';

COMMENT ON COLUMN "blogs"."updated_at" IS '修改时间';

ALTER TABLE "blogs" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");