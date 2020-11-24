package gnb

import (
	"fmt"
	"free5gc-cli/logger"
	"free5gc-cli/module/gnb/api"
	"strings"
)

func executorConfiguration(in string) {
	s := strings.TrimSpace(in)
	if s == "configuration reload" {
		Reload()
		return
	}
}

func executorUE(in string) {

	cmd := strings.Split(strings.TrimSpace(in), " ")
	l := len(cmd)

	if l < 2 {
		return
	}

	first := cmd[1]
	if first == "register" {
		if l < 4 {
			return
		}
		u := cmd[3]
		ueInfo := strings.Split(u, "/")
		ue, err := api.Registration(ueInfo[0], ueInfo[1])
		if err != nil {
			logger.GNBLog.Errorln(fmt.Sprintf("Error registering supi %s on plmn %s", ueInfo[0], ueInfo[1]))
			return
		}
		l := append(*gnb.UE, *ue)
		gnb.UE = &l
		return
	}

}

// Executor parse CLI
func Executor(in string) {

	if strings.HasPrefix(in, "configuration") {
		executorConfiguration(in)
	}

	if strings.HasPrefix(in, "ue") {
		executorUE(in)
	}

	if strings.HasPrefix(in, "pdu") {
	}

	if strings.HasPrefix(in, "qos") {
	}

}
