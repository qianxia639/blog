package config

type Http struct {
	Port int    `mapstructure:"port"`
	Host string `mapstructure:"host"`
}
