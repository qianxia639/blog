package utils

import (
	"io/ioutil"

	"github.com/qianxia/blog/model"
)

func DeCode() ([]byte, *model.Config) {
	fileName := "./config/application.yaml"
	y := new(model.Config)
	yamlFile, _ := ioutil.ReadFile(fileName)

	return yamlFile, y
}
