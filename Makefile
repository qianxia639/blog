DB_URL=postgresql://root:root@localhost:5432/blog?sslmode=disable&timeZone=Asia/Shanghai

server:
	go run main.go

test:
	go test -v -cover ./...

postgres:
	sudo docker run -d -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root --restart=always --name postgres postgres:14-alpine

createdb:
	sudo docker exec -it postgres createdb --username=root --owner=root blog

redis:
	sudo docker run -d -p 6379:6379 --restart=always --name redis redis:6-alpine

migrateup:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migrateup1:
	migrate -path db/migration -database "${DB_URL}" -verbose up 1

migratedown:
	migrate -path db/migration -database "${DB_URL}" -verbose down

migratedown1:
	migrate -path db/migration -database "${DB_URL}" -verbose down 1

sqlc:
	sqlc generate

.PHONY: server test postgres createdb redis migrateup migrateup1 migratedown migratedown1 sqlc