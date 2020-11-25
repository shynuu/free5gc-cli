package gnb

import (
	"fmt"
	"free5gc-cli/module/gnb/api"
	"net"
	"strings"

	"github.com/c-bata/go-prompt"
)

var gnb *GNB

type UEI struct {
	UE *api.RanUeContext
}

type PDUSession struct {
	SessionID uint8
	Ipv4      string
	TEID      uint32
	DNN       string
	Snssai    Snssai
}

type GNB struct {
	UE          *[]UEI
	PDUSessions *[]PDUSession
	SessionID   uint8
	Ipv4        net.IP
}

func (g *GNB) AlreadyRegister(supi string) bool {
	for _, ue := range *g.UE {
		if supi == ue.UE.Supi {
			return true
		}
	}
	return false
}

func (g *GNB) AddUE(ue *api.RanUeContext) {
	if g.AlreadyRegister(ue.Supi) {
		return
	}
	var uei = UEI{UE: ue}
	tmp := append(*g.UE, uei)
	g.UE = &tmp
	text := fmt.Sprintf("%s", ue.Supi)
	l := append(*RegisteredSuggestion, prompt.Suggest{Text: text, Description: ""})
	RegisteredSuggestion = &l
}

func NewGNB() *GNB {
	tmp := strings.Split(GNBConfig.Configuration.UESubnet, "/")
	var gnb = GNB{Ipv4: net.ParseIP(tmp[0])}
	gnb.UE = &[]UEI{}
	return &gnb
}
