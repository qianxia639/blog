package config

import "github.com/spf13/viper"

func LoadConfig(path string) (*Config, error) {

	var conf Config

	viper.AddConfigPath(path)

	viper.SetConfigName("config")
	viper.SetConfigType("toml")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&conf)
	return &conf, err
}
