package qos

import (
	"fmt"
	"free5gc-cli/logger"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var QOSConfig *Config

func checkErr(err error) {
	if err != nil {
		err = fmt.Errorf("[Configuration] %s", err.Error())
		logger.QOSLog.Fatal(err)
	}
}

// InitConfigFactory initialize the module configuration
func InitConfigFactory(f string, force bool) {

	if !force && QOSConfig != nil {
		return
	}

	content, err := ioutil.ReadFile(f)
	checkErr(err)

	QOSConfig = &Config{}

	err = yaml.Unmarshal([]byte(content), &QOSConfig)
	checkErr(err)
	logger.QOSLog.Infof("Successfully load module configuration %s", f)

}
