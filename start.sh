#!/bin/sh

set -e

# SOURCE=postgresql://root:root@postgres:5432/blog?sslmode=disable

# echo "${POSTGRES.SOURCE}"
echo "run db migration"
/app/migrate -path /app/migration -database "${POSTGRES.SOURCE}" -verbose up
echo "start the app"
exec "$@"