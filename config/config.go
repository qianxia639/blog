package config

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	}
	MySQL struct {
		Port     int    `yaml:"port"`
		Host     string `yaml:"host"`
		DbName   string `yaml:"dbname"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Charset  string `yaml:"charset"`
		Loc      string `yaml:"loc"`
	}
	Qiniu struct {
		AccessKey string `yaml:"access_key"`
		SecretKey string `yaml:"secret_key"`
		Bucket    string `yaml:"bucket"`
		ServerUrl string `yaml:"server_url"`
	}
}
