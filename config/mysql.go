package config

type MySQL struct {
	MaxIdle  int    `mapstructure:"max_idle"`
	MaxOpen  int    `mapstructure:"max_open"`
	Port     int    `mapstructure:"port"`
	Host     string `mapstructure:"host"`
	DbName   string `mapstructure:"dbname"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	Charset  string `mapstructure:"charset"`
	Loc      string `mapstructure:"loc"`
}
