package config

import "time"

type Config struct {
	Postgres Postgres `mapstructure:"postgres"`
	Token    Token    `mapstructure:"token"`
	Server   Server   `mapstructure:"server"`
	Redis    Redis    `mapstructure:"redis"`
	OssQiniu OssQiniu `mapstructure:"oss_qiniu"`
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

type OssQiniu struct {
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	ServerUrl string `mapstructure:"server_url"`
	Bucket    string `mapstructure:"bucket"`
}
