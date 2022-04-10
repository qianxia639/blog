package config

type Config struct {
	Http struct {
		Port int    `mapstructure:"port" yaml:"port" toml:"port"`
		Host string `mapstructure:"host" yaml:"host" toml:"host"`
	} `mapstructure:"http"`
	MySQL struct {
		MaxIdle  int    `mapstructure:"max_idle" yaml:"max_idle" toml:"max_idle"`
		MaxOpen  int    `mapstructure:"max_open" yaml:"max_open" toml:"max_open"`
		Port     int    `mapstructure:"port" yaml:"port" toml:"port"`
		Host     string `mapstructure:"host" yaml:"host" toml:"host"`
		DbName   string `mapstructure:"dbname" yaml:"dbname" toml:"dbname"`
		Username string `mapstructure:"username" yaml:"username" toml:"username"`
		Password string `mapstructure:"password" yaml:"password" toml:"password"`
		Charset  string `mapstructure:"charset" yaml:"charset" toml:"charset"`
	} `mapstructure:"mysql"`
	Qiniu struct {
		AccessKey string `mapstructure:"access_key" yaml:"access_key" toml:"access_key"`
		SecretKey string `mapstructure:"secret_key" yaml:"secret_key" toml:"secret_key"`
		Bucket    string `mapstructure:"bucket" yaml:"bucket" toml:"bucket"`
		ServerUrl string `mapstructure:"server_url" yaml:"server_url" toml:"server_url"`
	} `mapstructure:"qiniu"`
}
