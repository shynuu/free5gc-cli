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
		s = append(s, prompt.Suggest{Text: fmt.Sprintf("%s/%s", ue.Supi, ue.PLMN), Description: ""})
	}
	userSuggestion = &s

	var l []prompt.Suggest
	for _, sn := range *GNBConfig.Configuration.Snssai {
		snssai := fmt.Sprintf("%02d%s", sn.Sst, sn.Sd)
		l = append(l, prompt.Suggest{Text: snssai, Description: ""})
	}
	snssaiSuggestion = &l

	gnb = &GNB{}

}

// Reload the module
func Reload() {
	DefaultCLIConfigPath := "config/" + MODULE_NAME + ".yaml"
	InitConfigFactory(DefaultCLIConfigPath, true)
	api.Initialize(DefaultCLIConfigPath, true)

	var s []prompt.Suggest
	for _, ue := range *GNBConfig.Configuration.UEList {
		s = append(s, prompt.Suggest{Text: fmt.Sprintf("%s/%s", ue.Supi, ue.PLMN), Description: ""})
	}
	userSuggestion = &s

	var l []prompt.Suggest
	for _, sn := range *GNBConfig.Configuration.Snssai {
		snssai := fmt.Sprintf("%02d%s", sn.Sst, sn.Sd)
		l = append(l, prompt.Suggest{Text: snssai, Description: ""})
	}
	snssaiSuggestion = &l

}

// Exit and free the resources used by the module
func Exit() {}
