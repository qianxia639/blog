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

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/hibiken/asynq"
	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func main() {
	defer logs.Logger.Sync()
	conf, err := config.LoadConfig(".")
	if err != nil {
		logs.Logger.Fatal("load config err: ", zap.Error(err))
	}

	conn, err := sql.Open(conf.Postgres.Driver, conf.Postgres.Source)
	if err != nil {
		logs.Logger.Fatal("can't connect to db: ", zap.Error(err))
	}

	runDBMigrate(conf.Postgres.MigrateUrl, conf.Postgres.Source)

	store := db.NewStore(conn)

	runGinServer(conf, store)
}

func runDBMigrate(migrateUrl, dbSource string) {
	migration, err := migrate.New(migrateUrl, dbSource)
	if err != nil {
		logs.Logger.Error("can't create new migrate instance: ", zap.Error(err))
	}

	if err := migration.Up(); err != nil && err != migrate.ErrNoChange {
		logs.Logger.Error("failed to run migrate up: ", zap.Error(err))
	}

	logs.Logger.Info("db migrated successfully")
}

func runTaskProcessor(redisOpt asynq.RedisClientOpt) {
	taskProcessor := task.NewRedisTaskProcessor(redisOpt)

	err := taskProcessor.Start()
	if err != nil {
		logs.Logger.Fatal("failed to start task processor", zap.Error(err))
	}

	logs.Logger.Info("start task processor")
}

func runGinServer(conf *config.Config, store db.Store) {

	rdb := cache.InitRedis(conf)

	maker, err := token.NewPasetoMaker(conf.Token.TokenSymmetricKey)
	if err != nil {
		logs.Logger.Fatal("failed to create paseto maker", zap.Error(err))
	}

	redisOpt := asynq.RedisClientOpt{
		Addr: conf.Redis.Address,
	}

	taskDistributor := task.NewRedisTaskDistributor(redisOpt)
	go runTaskProcessor(redisOpt)

	opts := []api.ServerOptions{
		api.Conf(conf),
		api.Store(store),
		api.Cache(rdb),
		api.Maker(maker),
		api.TaskDistributor(taskDistributor),
		api.Router(),
	}

	server := api.NewServer(opts...)

	// log.Fatal(server.Start(conf.Server.Address))
	logs.Logger.Fatal("failed server start", zap.Error(server.Start(conf.Server.Address)))
}
