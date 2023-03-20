# Build
FROM golang:alpine AS builder

WORKDIR /app

COPY . .

ENV GOPROXY=https://goproxy.cn

RUN go build -o blog main.go
RUN apk update && \
    apk add --no-cache curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz 

# RUN
FROM golang:alpine

WORKDIR /app

COPY --from=builder /app/blog .
COPY --from=builder /app/migrate ./migrate
COPY config.toml .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 9010

CMD ["/app/blog"]
ENTRYPOINT ["/app/start.sh"]