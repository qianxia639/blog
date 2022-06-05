package config

type Config struct {
	Http struct {
		Port int    `mapstructure:"port"`
		Host string `mapstructure:"host"`
	}
	MySQL struct {
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
	Redis struct {
		Port     int    `mapstructure:"port"`
		DB       int    `mapstructure:"db"`
		Host     string `mapstructure:"host"`
		Password string `mapstructure:"password"`
	}
	Qiniu struct {
		AccessKey string `mapstructure:"access_key"`
		SecretKey string `mapstructure:"secret_key"`
		Bucket    string `mapstructure:"bucket"`
		ServerUrl string `mapstructure:"server_url"`
	}
	Captcha struct {
		Height   int     `mapstructure:"height"`
		Width    int     `mapstructure:"width"`
		Length   int     `mapstructure:"length"`
		MaxSkew  float64 `mapstructure:"max_skew"`
		DotCount int     `mapstructure:"dot_count"`
	}
}
