CREATE TABLE "blogs" (
  "id" bigserial PRIMARY KEY,
  "owner_id" int8 NOT NULL,
  "type_id" int8 NOT NULL,
  "title" varchar(50) NOT NULL UNIQUE,
  "content" text NOT NULL,
  "image" varchar(255) NOT NULL,
  "views" int4 DEFAULT 0 NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "types" (
  "id" bigserial PRIMARY KEY,
  "type_name" varchar(20) UNIQUE NOT NULL
);

COMMENT ON COLUMN "blogs"."id" IS '主键';

COMMENT ON COLUMN "blogs"."owner_id" IS '创建者Id';

COMMENT ON COLUMN "blogs"."type_id" IS '分类Id';

COMMENT ON COLUMN "blogs"."title" IS '标题';

COMMENT ON COLUMN "blogs"."content" IS '内容';

COMMENT ON COLUMN "blogs"."image" IS '图片链接';

COMMENT ON COLUMN "blogs"."views" IS '浏览次数';

COMMENT ON COLUMN "blogs"."created_at" IS '创建时间';

COMMENT ON COLUMN "blogs"."updated_at" IS '修改时间';

COMMENT ON COLUMN "types"."id" IS '主键';

COMMENT ON COLUMN "types"."type_name" IS '类别';

ALTER TABLE "blogs" ADD FOREIGN KEY ("owner_id") REFERENCES "users" ("id");

ALTER TABLE "blogs" ADD FOREIGN KEY ("type_id") REFERENCES "types" ("id");