package main

import (
	"Blog/api"
	db "Blog/db/sqlc"
	"Blog/utils/config"
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
		log.Fatal("can't create new migratee instance: ", err)
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up: ", err)
	}

	log.Println("db migrated successfully")
}

func runGinServer(conf config.Config, store db.Store) {
	server, err := api.NewServer(conf, store)
	if err != nil {
		log.Fatal("can't create serveere: ", err)
	}

	log.Fatal(server.Start(conf.Server.Address))
}
