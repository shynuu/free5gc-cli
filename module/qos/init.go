package qos

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

func Initialize() {
	DefaultCLIConfigPath := "config/" + MODULE_NAME + ".yaml"
	InitConfigFactory(DefaultCLIConfigPath, false)

	var l = []prompt.Suggest{}
	for _, ip := range QOSConfig.Configuration.Ipv4 {
		l = append(l, prompt.Suggest{Text: ip, Description: ""})
	}
	IPSuggestion = &l

	var teidTmp = []prompt.Suggest{}
	for i := 0; i < MAX_TEID; i++ {
		teidTmp = append(teidTmp, prompt.Suggest{Text: fmt.Sprintf("%03d", i), Description: ""})
	}
	TEIDSuggestion = &teidTmp

	var portTmp = []prompt.Suggest{}
	for _, port := range QOSConfig.Configuration.Port {
		portTmp = append(portTmp, prompt.Suggest{Text: port, Description: ""})
	}
	PortSuggestion = &portTmp

}

func Reload() {
	DefaultCLIConfigPath := "config/" + MODULE_NAME + ".yaml"
	InitConfigFactory(DefaultCLIConfigPath, true)

	var l = []prompt.Suggest{}
	for _, ip := range QOSConfig.Configuration.Ipv4 {
		l = append(l, prompt.Suggest{Text: ip, Description: ""})
	}
	IPSuggestion = &l

	var teidTmp = []prompt.Suggest{}
	for i := 0; i < MAX_TEID; i++ {
		teidTmp = append(teidTmp, prompt.Suggest{Text: fmt.Sprintf("%03d", i), Description: ""})
	}
	TEIDSuggestion = &teidTmp

	var portTmp = []prompt.Suggest{}
	for _, port := range QOSConfig.Configuration.Port {
		portTmp = append(portTmp, prompt.Suggest{Text: port, Description: ""})
	}
	PortSuggestion = &portTmp
}

func Exit() {

}
