package utils

import (
	"io/ioutil"
	"os"

	"github.com/qianxia/blog/config"
	"github.com/qianxia/blog/global"
	"gopkg.in/yaml.v2"
)

func DeCode() (config *config.Config) {
	dir, _ := os.Getwd()
	fileName := dir + "/config/application.yaml"
	yamlFile, _ := ioutil.ReadFile(fileName)
	err := yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		global.QX_LOG.Errorf("%v", err)
		return nil
	}
	return
}
