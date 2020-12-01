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

	var q []prompt.Suggest
	for _, dnn := range GNBConfig.Configuration.DNN {
		dnna := fmt.Sprintf("%s", dnn)
		q = append(q, prompt.Suggest{Text: dnna, Description: ""})
	}
	DNNSuggestion = &q

	gnb = NewGNB()
	gtpRouter, err := NewRouter(GNBConfig.Configuration.UpfInterface.IPv4Addr,
		GNBConfig.Configuration.UpfInterface.Port,
		GNBConfig.Configuration.GTPInterface.Ipv4,
		GNBConfig.Configuration.GTPInterface.Port,
		GNBConfig.Configuration.UESubnet,
		gnb,
	)
	if err != nil {
		panic("Impossible to start the gNB router")
	}
	go gtpRouter.Encapsulate()
	go gtpRouter.Desencapsulate()

}

// Reload the module
func Reload() {

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

	var q []prompt.Suggest
	for _, dnn := range GNBConfig.Configuration.DNN {
		dnna := fmt.Sprintf("%s", dnn)
		q = append(l, prompt.Suggest{Text: dnna, Description: ""})
	}
	DNNSuggestion = &q

	gnb = NewGNB()
	gtpRouter.GNB = gnb

}

// Exit and free the resources used by the module
func Exit() {
	gtpRouter.Close()
}
