package nf

import (
	"fmt"
	"free5gc-cli/logger"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var NFConfig *Config

func checkErr(err error) {
	if err != nil {
		err = fmt.Errorf("[Configuration] %s", err.Error())
		logger.NFLog.Fatal(err)
	}
}

// InitConfigFactory initialize the module configuration
func InitConfigFactory(f string, force bool) {

	if !force && NFConfig != nil {
		return
	}

	content, err := ioutil.ReadFile(f)
	checkErr(err)

	NFConfig = &Config{}

	err = yaml.Unmarshal([]byte(content), &NFConfig)
	checkErr(err)
	logger.NFLog.Infof("Successfully load module configuration %s", f)
}
