package utils

import (
	"io/ioutil"
	"os"
	"runtime"

	"github.com/BurntSushi/toml"
	"github.com/qianxia/blog/config"
	"github.com/qianxia/blog/global"
)

func DeCodeTOML(path string) (config *config.Config) {

	switch runtime.GOOS {
	case "windows":
		dir, _ := os.Getwd()
		path = dir + "/config/application.toml"
	case "linux":
		path = "/opt/conf/application.toml"
	}

	tomlFile, _ := ioutil.ReadFile(path)
	err := toml.Unmarshal(tomlFile, &config)
	if err != nil {
		global.QX_LOG.Errorf("decode toml err: %v", err)
		return
	}
	return
}
