package config

type Qiniu struct {
	AccessKey string `mapstructure:"access_key" toml:"access_key"`
	SecretKey string `mapstructure:"secret_key" toml:"secret_key"`
	Bucket    string `mapstructure:"bucket" toml:"bucket"`
	ServerUrl string `mapstructure:"server_url" toml:"server_url"`
}
