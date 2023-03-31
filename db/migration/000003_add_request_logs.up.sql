CREATE TABLE "request_logs" (
  "id" bigserial PRIMARY KEY,
  "method" varchar(10) NOT NULL,
  "path" varchar(30) NOT NULL,
  "status_code" int4 NOT NULL,
  "ip" varchar(30) NOT NULL,
  "hostname" varchar(30) NOT NULL,
  "request_body" text NOT NULL DEFAULT '',
  "response_time" int8 NOT NULL DEFAULT 0,
  "request_time" timestamptz NOT NULL DEFAULT(now()),
  "content_type" varchar(40) NOT NULL DEFAULT '',
  "user_agent" varchar(100) NOT NULL
);

COMMENT ON COLUMN "request_logs"."id" IS '主键';

COMMENT ON COLUMN "request_logs"."method" IS '请求方式';

COMMENT ON COLUMN "request_logs"."path" IS '路由';

COMMENT ON COLUMN "request_logs"."status_code" IS '状态码';

COMMENT ON COLUMN "request_logs"."ip" IS '访问ip';

COMMENT ON COLUMN "request_logs"."hostname" IS '主机名';

COMMENT ON COLUMN "request_logs"."request_body" IS '请求体';

COMMENT ON COLUMN "request_logs"."response_time" IS '响应时间/ms';

COMMENT ON COLUMN "request_logs"."request_time" IS '请求时间';

COMMENT ON COLUMN "request_logs"."content_type" IS '请求数据类型';

COMMENT ON COLUMN "request_logs"."user_agent" IS 'user_agent';