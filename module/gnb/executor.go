package gnb

import (
	"fmt"
	"free5gc-cli/logger"
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

	if l < 2 || l > 4 {
		return
	}

	first := cmd[1]

	if first == "register" {
		if l < 4 {
			return
		}
		u := cmd[3]
		ueInfo := strings.Split(u, "/")
		logger.GNBLog.Infoln(fmt.Sprintf("Registering user %s on the network", ueInfo[0]))
		err := gnb.Register(ueInfo[0])
		if err != nil {
			return
		}
		logger.GNBLog.Infoln(fmt.Sprintf("Successfully register user %s on the network", ueInfo[0]))
		return
	}

	if first == "deregister" {
		if l < 4 {
			return
		}
		ue := cmd[3]
		logger.GNBLog.Infoln(fmt.Sprintf("De-Registering user %s on the network", ue))
		err := gnb.Deregister(ue)
		if err != nil {
			return
		}
		logger.GNBLog.Infoln(fmt.Sprintf("Successfully de-register user %s on the network", ue))
		return
	}

	return

}

func executorPDUSession(in string) {

	cmd := strings.Split(strings.TrimSpace(in), " ")
	l := len(cmd)

	if l < 2 {
		return
	}

	first := cmd[1]

	if first == "request" && l > 8 {
		logger.GNBLog.Infoln(fmt.Sprintf("Establishing PDU Session for user %s with snssai %s and dnn %s", cmd[3], cmd[5], cmd[7]))
		err := gnb.PDURequest(cmd[3], cmd[5], cmd[7])
		if err != nil {
			return
		}
		logger.GNBLog.Infoln(fmt.Sprintf("Successfully Established PDU Session for user %s with snssai %s and dnn %s", cmd[3], cmd[5], cmd[7]))
	}

	if first == "release" && l > 3 {
		cmd = strings.Split(cmd[3], "-")
		logger.GNBLog.Infoln(fmt.Sprintf("Releasing PDU Session for user %s with session %s", cmd[0], cmd[1]))
		err := gnb.PDURelease(cmd[0], cmd[1])
		if err != nil {
			return
		}
		logger.GNBLog.Infoln(fmt.Sprintf("Successfully Releasing PDU Session for user %s with session %s", cmd[0], cmd[1]))
	}

	return

}

// Executor parse CLI
func Executor(in string) {

	if strings.HasPrefix(in, "configuration") {
		executorConfiguration(in)
	}

	if strings.HasPrefix(in, "user") {
		executorUE(in)
	}

	if strings.HasPrefix(in, "pdu-session") {
		executorPDUSession(in)
	}

	if strings.HasPrefix(in, "qos") {
	}

}
