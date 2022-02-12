package config

type Config struct {
	Server struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	}
	MySQL struct {
		Port       int    `yaml:"port"`
		Host       string `yaml:"host"`
		DriverName string `yaml:"driver_name"`
		DbName     string `yaml:"dbname"`
		Username   string `yaml:"username"`
		Password   string `yaml:"password"`
		Charset    string `yaml:"charset"`
		Loc        string `yaml:"loc"`
	}
	Redis struct {
		Port int    `yaml:"port"`
		Host string `yaml:"host"`
	}
	QQMail struct {
		Port     int    `yaml:"port"`
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}
}
