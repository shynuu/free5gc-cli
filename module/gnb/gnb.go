package gnb

import "free5gc-cli/module/gnb/api"

var gnb *GNB

type GNB struct {
	UE *[]api.RanUeContext
}
