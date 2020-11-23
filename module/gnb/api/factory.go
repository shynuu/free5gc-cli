package api

import (
	"fmt"
	"free5gc-cli/logger"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var APIConfig *Config

func checkErr(err error) {
	if err != nil {
		err = fmt.Errorf("[Configuration] %s", err.Error())
		logger.GNBLog.Fatal(err)
	}
}

// InitConfigFactory initialize the module configuration
func InitConfigFactory(f string, force bool) {

	if !force && APIConfig != nil {
		return
	}

	content, err := ioutil.ReadFile(f)
	checkErr(err)

	APIConfig = &Config{}

	err = yaml.Unmarshal([]byte(content), &APIConfig)
	checkErr(err)
	logger.GNBLog.Infof("Successfully load gNB API module configuration %s", f)
}
