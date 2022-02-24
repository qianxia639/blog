package utils

import (
	"io/ioutil"

	"github.com/qianxia/blog/config"
	"github.com/qianxia/blog/global"
	"gopkg.in/yaml.v2"
)

func DeCode() (config *config.Config) {
	fileName := "./config/application.yaml"
	yamlFile, _ := ioutil.ReadFile(fileName)
	err := yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		global.RY_LOG.Errorf("%v", err)
		return nil
	}
	return
}
