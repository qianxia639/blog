package main

import (
	"Blog/api"
	"Blog/core/cache"
	"Blog/core/config"
	"Blog/core/logs"
	"Blog/core/token"
	db "Blog/db/sqlc"
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
)

func main() {

	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("load config err: ", err)
	}

	logs.Logs = logs.InitZap(&conf)

	conn, err := sql.Open(conf.Postgres.Driver, conf.Postgres.Source)
	if err != nil {
		log.Fatal("can't connect to db: ", err)
	}

	runDBMigrate(conf.Postgres.MigrateUrl, conf.Postgres.Source)

	store := db.NewStore(conn)

	runGinServer(conf, store)
}

func runDBMigrate(migrateUrl, dbSource string) {
	migration, err := migrate.New(migrateUrl, dbSource)
	if err != nil {
		logs.Logs.Error("can't create new migrate instance: ", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		logs.Logs.Error("failed to run migrate up: ", err)
	}

	logs.Logs.Info("db migrated successfully")
}

func runGinServer(conf config.Config, store db.Store) {

	rdb := cache.InitRedis(conf)

	maker, err := token.NewPasetoMaker(conf.Token.TokenSymmetricKey)
	if err != nil {
		log.Fatal(err)
	}

	opts := []api.ServerOptions{
		api.WithStor(store),
		api.WithConfig(conf),
		api.WithMaker(maker),
		api.WithCache(rdb),
	}

	server := api.NewServer(opts...)

	log.Fatal(server.Start(conf.Server.Address))
}
