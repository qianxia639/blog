package config

import "time"

type Config struct {
	Postgres Postgres
	Token    Token
	Server   Server
	Redis    Redis
	Zap      Zap
}

type Postgres struct {
	Driver     string `mapstructure:"driver"`
	Source     string `mapstructure:"source"`
	MigrateUrl string `mapstructure:"migrate_url"`
}

type Token struct {
	TokenSymmetricKey   string        `mapstructure:"token_symmetric_key"`
	AccessTokenDuration time.Duration `mapstructure:"access_token_duration"`
}

type Server struct {
	Address string `mapstructure:"address"`
}

type Redis struct {
	Address  string `mapstructure:"address"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type Zap struct {
	Prefix     string    `mapstructure:"prefix"`
	TimeFormat time.Time `mapstructure:"time_format"`
	Level      string    `mapstructure:"level"`
	Caller     bool      `mapstructure:"caller"`
	StackTrace bool      `mapstructure:"stack_trace"`
	Writer     string    `mapstructure:"writer"`
	Encode     string    `mapstructure:"encode"`
	LogFile    *LogFile  `mapstructure:"log_file"`
}

type LogFile struct {
	MaxSize  int    `mapstructure:"max_size"`
	Backups  int    `mapstructure:"backups"`
	Compress bool   `mapstructure:"compress"`
	Output   string `mapstructure:"output"`
}
