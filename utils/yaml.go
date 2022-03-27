package utils

import (
	"io/ioutil"
	"os"
	"runtime"

	"github.com/qianxia/blog/config"
	"github.com/qianxia/blog/global"
	"gopkg.in/yaml.v2"
)

func ParseConfig() (config *config.Config) {
	var fileName string
	if runtime.GOOS == "windows" {
		dir, _ := os.Getwd()
		fileName = dir + "/config/application.yaml"
	} else if runtime.GOOS == "linux" {
		fileName = "/opt/conf/application.yaml"
	}

	yamlFile, _ := ioutil.ReadFile(fileName)
	err := yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		global.QX_LOG.Errorf("%v", err)
		return nil
	}
	return
}
