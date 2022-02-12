package utils

import (
	"io/ioutil"

	"github.com/qianxia/blog/config"
)

func DeCode() ([]byte, *config.Config) {
	fileName := "./config/application.yaml"
	yamlFile, _ := ioutil.ReadFile(fileName)
	return yamlFile, &config.Config{}
}
