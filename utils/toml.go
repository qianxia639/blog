package utils

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/qianxia/blog/config"
	"github.com/qianxia/blog/global"
)

func DeCode() (config *config.Config) {
	fileName := "./config/application.toml"
	tomlFile, _ := ioutil.ReadFile(fileName)
	err := toml.Unmarshal(tomlFile, &config)
	if err != nil {
		global.QX_LOG.Errorf("decode toml err: %v", err)
		return
	}
	return
}
