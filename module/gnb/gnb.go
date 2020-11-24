package gnb

import "free5gc-cli/module/gnb/api"

var gnb *GNB

type GNB struct {
	UE *[]api.RanUeContext
}

func NewGNB() *GNB {
	var gnb = GNB{}
	l := make([]api.RanUeContext, 1)
	gnb.UE = &l
	return &gnb
}
