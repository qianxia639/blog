DB_URL=mysql://root:qianxia@tcp(localhost:3306)/blog?multiStatements=true

migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

sqlc:
	sqlc generate

.PHONY: migrateup migratedown sqlc