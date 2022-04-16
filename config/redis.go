package config

type Redis struct {
	Port     int    `mapstructure:"port"`
	DB       int    `mapstructure:"db"`
	Host     string `mapstructure:"host"`
	Password string `mapstructure:"password"`
}
