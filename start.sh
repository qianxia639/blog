#!/bin/sh

set -e

source="postgresql://root:root@postgres:5432/blog?sslmode=disable&timeZone=Asia/Shanghai"

echo "run db migration"
/app/migrate -path /app/migration -database "${source}" -verbose up
echo "start the app"
exec "$@"