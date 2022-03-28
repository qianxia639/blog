package config

type Config struct {
	Server struct {
		Port int    `yaml:"port" toml:"port"`
		Host string `yaml:"host" toml:"host"`
	}
	MySQL struct {
		MaxIdle  int    `yaml:"max_idle" toml:"max_idle"`
		MaxOpen  int    `yaml:"max_open" toml:"max_open"`
		Port     int    `yaml:"port" toml:"port"`
		Host     string `yaml:"host" toml:"host"`
		DbName   string `yaml:"dbname" toml:"dbname"`
		Username string `yaml:"username" toml:"username"`
		Password string `yaml:"password" toml:"password"`
		Charset  string `yaml:"charset" toml:"charset"`
	}
	Qiniu struct {
		AccessKey string `yaml:"access_key" toml:"access_key"`
		SecretKey string `yaml:"secret_key" toml:"secret_key"`
		Bucket    string `yaml:"bucket" toml:"bucket"`
		ServerUrl string `yaml:"server_url" toml:"server_url"`
	}
}
