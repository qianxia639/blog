package config

type Qiniu struct {
	AccessKey string `mapstructure:"access_key"`
	SecretKey string `mapstructure:"secret_key"`
	Bucket    string `mapstructure:"bucket"`
	ServerUrl string `mapstructure:"server_url"`
}
