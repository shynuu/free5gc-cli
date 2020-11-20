package subscriber

import (
	"fmt"
	"free5gc-cli/logger"
	"free5gc-cli/module/subscriber/api"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

var SubscriberConfig *Config
var SubsDataConfig *api.SubsData

func checkErr(err error) {
	if err != nil {
		err = fmt.Errorf("[Configuration] %s", err.Error())
		logger.SubscriberLog.Fatal(err)
	}
}

// InitConfigFactory initialize the module configuration
func InitConfigFactory(f string, force bool) {

	if !force && SubscriberConfig != nil {
		return
	}

	content, err := ioutil.ReadFile(f)
	checkErr(err)

	SubscriberConfig = &Config{}

	err = yaml.Unmarshal([]byte(content), &SubscriberConfig)
	checkErr(err)
	logger.SubscriberLog.Infof("Successfully load module configuration %s", f)
}

// InitializeUEConfiguration initialize the ue configuration
func InitializeUEConfiguration(f string, force bool) {

	if !force && SubsDataConfig != nil {
		return
	}

	content, err := ioutil.ReadFile(f)
	checkErr(err)

	SubsDataConfig = &api.SubsData{}

	err = yaml.Unmarshal([]byte(content), &SubsDataConfig)
	checkErr(err)
	logger.SubscriberLog.Infof("Successfully load ue configuration %s", f)

}
