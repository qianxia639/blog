package model

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	}
	MySQL struct {
		DriverName string `yaml:"driver_name"`
		Host       string `yaml:"host"`
		Port       int    `yaml:"port"`
		DbName     string `yaml:"dbname"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Charset    string `yaml:"charset"`
		Loc        string `yaml:"loc"`
	}
	Redis struct {
		Host string `yaml:"host"`
		Port int    `yaml:"port"`
	}
	QQMail struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Port     int    `yaml:"port"`
		Host     string `yaml:"host"`
	}
}
