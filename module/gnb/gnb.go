package gnb

import (
	"errors"
	"fmt"
	"free5gc-cli/logger"
	"free5gc-cli/module/gnb/api"
	"net"
	"strconv"
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

func removeFromPDUSession(s []PDUSession, index int, length int) []PDUSession {
	if length == 1 {
		return []PDUSession{}
	}
	if index == length-1 {
		return append(s[:index-1])
	}
	return append(s[:index], s[index+1:]...)
}

var gnb *GNB

type PDUSession struct {
	Supi      string
	SessionID uint8
	Ipv4      net.IP
	TEID      uint32
	DNN       string
	Snssai    string
}

type GNB struct {
	UE          *[]api.RanUeContext
	PDUSessions *[]PDUSession
	SessionID   uint8
	Ipv4        *net.IP
	TEID        uint32
}

func (g *GNB) IncrementIP() {
	var b3 byte = (*g.Ipv4)[3] + 1
	Ipv4 := net.IPv4((*g.Ipv4)[0], (*g.Ipv4)[1], (*g.Ipv4)[2], b3)
	g.Ipv4 = &Ipv4
}

func (g *GNB) DecrementIP() {
	var b3 byte = (*g.Ipv4)[3] - 1
	Ipv4 := net.IPv4((*g.Ipv4)[0], (*g.Ipv4)[1], (*g.Ipv4)[2], b3)
	g.Ipv4 = &Ipv4
}

func (g *GNB) IncrementSessionId() {
	g.SessionID++
}

func (g *GNB) DecrementSessionId() {
	g.SessionID--
}

func (g *GNB) IncrementTEID() {
	g.TEID++
}

func (g *GNB) DecrementTEID() {
	g.TEID--
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

func (g *GNB) GetUEContext(supi string) (*api.RanUeContext, error) {

	for _, ue := range *g.UE {
		if ue.Supi == supi {
			return &ue, nil
		}
	}
	return nil, errors.New("")
}

func (g *GNB) AddPDU(pdu PDUSession) {
	tmp := append(*g.PDUSessions, pdu)
	g.PDUSessions = &tmp

	l := append(*PDUSuggestion, prompt.Suggest{Text: fmt.Sprintf("%s-%d", pdu.Supi, pdu.SessionID)})
	PDUSuggestion = &l
}

func (g *GNB) DeletePDU(pdu PDUSession) {

	for i, pduu := range *g.PDUSessions {
		if pdu.Supi == pduu.Supi && pdu.SessionID == pduu.SessionID {
			l := removeFromPDUSession(*g.PDUSessions, i, len(*g.PDUSessions))
			g.PDUSessions = &l
			break
		}
	}

	for i, pduu := range *PDUSuggestion {
		if pduu.Text == fmt.Sprintf("%s-%d", pdu.Supi, pdu.SessionID) {
			l := removeIndex(*PDUSuggestion, i, len(*PDUSuggestion))
			PDUSuggestion = &l
			break
		}
	}

}

func (g *GNB) PDURequest(supi string, snssai string, dnn string) error {

	ue, err := g.GetUEContext(supi)
	if err != nil {
		logger.GNBLog.Errorln(fmt.Sprintf("Impossible to find in UE %s", supi))
		return err
	}
	g.IncrementSessionId()
	err = api.PDUSessionRequest(ue, snssai, g.SessionID, dnn)
	if err != nil {
		logger.GNBLog.Errorln(fmt.Sprintf("Impossible to establish a PDU session for UE %s", supi))
		g.DecrementSessionId()
		return err
	}
	g.IncrementIP()
	g.IncrementTEID()

	pdu := PDUSession{SessionID: g.SessionID, Snssai: snssai, TEID: g.TEID, Ipv4: *g.Ipv4, Supi: supi}
	g.AddPDU(pdu)

	return nil
}

func (g *GNB) GetPDUSession(supi string, sessionID uint8) (*PDUSession, error) {

	for _, pdu := range *g.PDUSessions {
		if pdu.Supi == supi && pdu.SessionID == sessionID {
			return &pdu, nil
		}
	}

	return nil, errors.New("")

}

func (g *GNB) PDURelease(supi string, sessionid string) error {

	ue, err := g.GetUEContext(supi)
	if err != nil {
		logger.GNBLog.Errorln(fmt.Sprintf("Impossible to find in UE %s", supi))
		return err
	}

	sessionID64, err := strconv.Atoi(sessionid)
	sessionID := uint8(sessionID64)

	pdu, err := g.GetPDUSession(supi, sessionID)
	if err != nil {
		logger.GNBLog.Errorln(fmt.Sprintf("Impossible to find the PDU Session %d of user %s", sessionID, supi))
		return err
	}

	err = api.PDUSessionRelease(ue, pdu.Snssai, pdu.SessionID, pdu.DNN)
	if err != nil {
		logger.GNBLog.Errorln(fmt.Sprintf("Impossible to release the PDU Session %d of user %s", sessionID, supi))
		return err
	}

	g.DeletePDU(*pdu)

	return nil
}

func NewGNB() *GNB {
	tmp := strings.Split(GNBConfig.Configuration.UESubnet, "/")
	ipv4 := net.ParseIP(tmp[0]).To4()
	var gnb = GNB{Ipv4: &ipv4, SessionID: 9}
	gnb.UE = &[]api.RanUeContext{}
	return &gnb
}
