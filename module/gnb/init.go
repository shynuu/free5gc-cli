package gnb

import (
	"fmt"
	"free5gc-cli/module/gnb/api"

	"github.com/c-bata/go-prompt"
)

// Initialize the module
func Initialize() {
	DefaultCLIConfigPath := "config/" + MODULE_NAME + ".yaml"
	InitConfigFactory(DefaultCLIConfigPath, false)
	api.Initialize(DefaultCLIConfigPath, false)

	var s []prompt.Suggest
	for _, ue := range *GNBConfig.Configuration.UEList {
		s = append(s, prompt.Suggest{Text: fmt.Sprintf("%s", ue.Supi), Description: ""})
	}
	UserSuggestion = &s

	var l []prompt.Suggest
	for _, sn := range *GNBConfig.Configuration.Snssai {
		snssai := fmt.Sprintf("%02d%s", sn.Sst, sn.Sd)
		l = append(l, prompt.Suggest{Text: snssai, Description: ""})
	}
	SnssaiSuggestion = &l

	gnb = NewGNB()

}

// Reload the module
func Reload() {
	DefaultCLIConfigPath := "config/" + MODULE_NAME + ".yaml"
	InitConfigFactory(DefaultCLIConfigPath, true)
	api.Initialize(DefaultCLIConfigPath, true)

	var s []prompt.Suggest
	for _, ue := range *GNBConfig.Configuration.UEList {
		s = append(s, prompt.Suggest{Text: fmt.Sprintf("%s", ue.Supi), Description: ""})
	}
	UserSuggestion = &s
	var l []prompt.Suggest
	for _, sn := range *GNBConfig.Configuration.Snssai {
		snssai := fmt.Sprintf("%02d%s", sn.Sst, sn.Sd)
		l = append(l, prompt.Suggest{Text: snssai, Description: ""})
	}
	SnssaiSuggestion = &l

}

// Exit and free the resources used by the module
func Exit() {}
