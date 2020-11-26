package gnb

import (
	"fmt"
	"free5gc-cli/logger"
	"free5gc-cli/module/gnb/api"
	"net"
	"strings"

	"github.com/c-bata/go-prompt"
)

func removeIndex(s []prompt.Suggest, index int, length int) []prompt.Suggest {
	if length == 1 {
		return []prompt.Suggest{}
	}
	if index == length-1 {
		return append(s[:index-1])
	}
	return append(s[:index], s[index+1:]...)
}

func removeFromUE(s []api.RanUeContext, index int, length int) []api.RanUeContext {
	if length == 1 {
		return []api.RanUeContext{}
	}
	if index == length-1 {
		return append(s[:index-1])
	}
	return append(s[:index], s[index+1:]...)
}

var gnb *GNB

type PDUSession struct {
	SessionID uint8
	Ipv4      string
	TEID      uint32
	DNN       string
	Snssai    Snssai
}

type GNB struct {
	UE          *[]api.RanUeContext
	PDUSessions *[]PDUSession
	SessionID   uint8
	Ipv4        net.IP
}

func (g *GNB) AlreadyRegister(supi string) bool {
	for _, ue := range *g.UE {
		if supi == ue.Supi {
			return true
		}
	}
	return false
}

func (g *GNB) AddUE(ue *api.RanUeContext) {
	if g.AlreadyRegister(ue.Supi) {
		return
	}
	tmp := append(*g.UE, *ue)
	g.UE = &tmp
	text := fmt.Sprintf("%s", ue.Supi)
	l := append(*RegisteredSuggestion, prompt.Suggest{Text: text, Description: ""})
	RegisteredSuggestion = &l
}

func (g *GNB) Register(supi string) error {
	if gnb.AlreadyRegister(supi) {
		logger.GNBLog.Infoln(fmt.Sprintf("Supi %s already registered on the network", supi))
		return nil
	}
	ue, err := api.Registration(supi)
	if err != nil {
		logger.GNBLog.Errorln(fmt.Sprintf("Error registering supi %s", supi))
		return err
	}
	gnb.AddUE(ue)
	return nil
}

func (g *GNB) Deregister(supi string) error {
	for j, ue := range *g.UE {
		if ue.Supi == supi {
			err := api.DeRegistration(&ue)
			if err != nil {
				logger.GNBLog.Errorln(fmt.Sprintf("Error for unregister user with supi %s", supi))
				return err
			}
			ue := removeFromUE(*g.UE, j, len(*g.UE))
			g.UE = &ue

			for i, sugg := range *RegisteredSuggestion {
				if sugg.Text == supi {
					l := removeIndex(*RegisteredSuggestion, i, len(*RegisteredSuggestion))
					RegisteredSuggestion = &l
					return nil
				}
			}
		}
	}
	logger.GNBLog.Infoln(fmt.Sprintf("No registered user with supi %s", supi))
	return nil
}

func NewGNB() *GNB {
	tmp := strings.Split(GNBConfig.Configuration.UESubnet, "/")
	var gnb = GNB{Ipv4: net.ParseIP(tmp[0])}
	gnb.UE = &[]api.RanUeContext{}
	return &gnb
}
