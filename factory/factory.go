/*
 * Freecli Configuration Factory
 */

package factory

import (
	"fmt"
	"free5gc-cli/logger"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var FreecliConfig Config

func checkErr(err error) {
	if err != nil {
		err = fmt.Errorf("[Configuration] %s", err.Error())
		logger.AppLog.Fatal(err)
	}
}

// TODO: Support configuration update from REST api
func InitConfigFactory(f string) {
	content, err := ioutil.ReadFile(f)
	checkErr(err)

	FreecliConfig = Config{}

	err = yaml.Unmarshal([]byte(content), &FreecliConfig)
	checkErr(err)

	logger.InitLog.Infof("Successfully initialize configuration %s", f)
}
