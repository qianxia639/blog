package config

type MySQL struct {
	MaxIdle  int    `mapstructure:"max_idle" toml:"max_idle"`
	MaxOpen  int    `mapstructure:"max_open" toml:"max_open"`
	Port     int    `mapstructure:"port" toml:"port"`
	Host     string `mapstructure:"host" toml:"host"`
	DbName   string `mapstructure:"dbname" toml:"dbname"`
	Username string `mapstructure:"username" toml:"username"`
	Password string `mapstructure:"password" toml:"password"`
	Charset  string `mapstructure:"charset" toml:"charset"`
	Loc      string `mapstructure:"loc" toml:"loc"`
}
