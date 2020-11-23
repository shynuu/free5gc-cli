package gnb

import (
	"fmt"
	"free5gc-cli/module/gnb/api"

	"github.com/c-bata/go-prompt"
)

// ue register --supi imsi-20893XXXXXX00 --plmn popo
// ue deregister --supi imsi-20893XXXXXX00 --plmn popo
// ue list

// pdu-session list
// pdu-session request --supi POPOD --plmn 20893
// pdu-session
// ===> ipv4, qos profile, sessionid
// pdu-session release --session <session_id>
// pdu-session

// qos add --session 10 --protocol tcp --destination-port 80 --phb

// Initialize the module
func Initialize() {
	DefaultCLIConfigPath := "config/" + MODULE_NAME + ".yaml"
	InitConfigFactory(DefaultCLIConfigPath, false)
	api.Initialize(DefaultCLIConfigPath)

	var s []prompt.Suggest
	var p []prompt.Suggest
	for _, ue := range *GNBConfig.Configuration.UEList {
		s = append(s, prompt.Suggest{Text: ue.Supi, Description: ""})
		p = append(p, prompt.Suggest{Text: ue.PLMN, Description: ""})
	}
	supiSuggestion = &s
	plmnSuggestion = &p

	var l []prompt.Suggest
	for _, sn := range *GNBConfig.Configuration.Snssai {
		snssai := fmt.Sprintf("%02d%s", sn.Sst, sn.Sd)
		l = append(l, prompt.Suggest{Text: snssai, Description: ""})
	}
	snssaiSuggestion = &l

}
