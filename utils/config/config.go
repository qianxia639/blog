package config

import "time"

type Config struct {
	Postgres struct {
		Driver     string `mapstructure:"driver"`
		Source     string `mapstructure:"source"`
		MigrateUrl string `mapstructure:"migrate_url"`
	}
	Token struct {
		TokenSymmetricKey   string        `mapstructure:"token_symmetric_key"`
		AccessTokenDuration time.Duration `mapstructure:"access_token_duration"`
	}
	Server struct {
		Address string `mapstructure:"address"`
	}
}
