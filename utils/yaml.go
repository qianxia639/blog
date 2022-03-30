package utils

import (
	"io/ioutil"
	"os"
	"runtime"

	"github.com/qianxia/blog/config"
	"github.com/qianxia/blog/global"
	"gopkg.in/yaml.v2"
)

func DeCodeYAML(path string) (config *config.Config) {
	switch runtime.GOOS {
	case "windows":
		dir, _ := os.Getwd()
		path = dir + "/config/application.yaml"
	case "linux":
		path = "/opt/conf/application.yaml"
	}

	yamlFile, _ := ioutil.ReadFile(path)
	err := yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		global.QX_LOG.Errorf("decode yaml err: %v", err)
		return nil
	}
	return
}
