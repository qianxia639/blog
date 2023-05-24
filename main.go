package main

import (
	"Blog/api"
	"Blog/core/cache"
	"Blog/core/config"
	"Blog/core/logs"
	"Blog/core/task"
	"Blog/core/token"
	db "Blog/db/sqlc"
	"database/sql"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {

	conf, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("load config err: ", err)
	}

	// logs.Logs = logs.InitZap(conf.Zap)
	logs.Logs = logs.GetInstance(conf.Zap)
	defer logs.Logs.Sync()

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
		logs.Logs.Error("can't create new migrate instance: ", zap.Error(err))
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		logs.Logs.Error("failed to run migrate up: ", zap.Error(err))
	}

	logs.Logs.Info("db migrated successfully")
}

func runTaskProcessor(redisOpt asynq.RedisClientOpt) {
	taskProcessor := task.NewRedisTaskProcessor(redisOpt)

	err := taskProcessor.Start()
	if err != nil {
		logs.Logs.Fatal("failed to start task processor", zap.Error(err))
	}

	logs.Logs.Info("start task processor")
}

func runGinServer(conf *config.Config, store db.Store) {

	rdb := cache.InitRedis(conf)

	maker, err := token.NewPasetoMaker(conf.Token.TokenSymmetricKey)
	if err != nil {
		log.Fatal(err)
	}

	redisOpt := asynq.RedisClientOpt{
		Addr: conf.Redis.Address,
	}

	taskDistributor := task.NewRedisTaskDistributor(redisOpt)
	go runTaskProcessor(redisOpt)

	opts := []api.ServerOptions{
		api.WithStor(store),
		api.WithConfig(conf),
		api.WithMaker(maker),
		api.WithCache(rdb),
		api.WithTaskDistributor(taskDistributor),
	}

	server := api.NewServer(opts...)

	log.Fatal(server.Start(conf.Server.Address))
}
