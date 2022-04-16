package config

type Http struct {
	Port int    `mapstructure:"port" toml:"port"`
	Host string `mapstructure:"host" toml:"host"`
}
