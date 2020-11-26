package gnb

import (
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
		err := gnb.Register(ueInfo[0])
		if err != nil {
			return
		}
		return
	}

	if first == "deregister" {
		if l < 4 {
			return
		}
		ue := cmd[3]
		err := gnb.Deregister(ue)
		if err != nil {
			return
		}
	}

}

// Executor parse CLI
func Executor(in string) {

	if strings.HasPrefix(in, "configuration") {
		executorConfiguration(in)
	}

	if strings.HasPrefix(in, "user") {
		executorUE(in)
	}

	if strings.HasPrefix(in, "pdu") {
	}

	if strings.HasPrefix(in, "qos") {
	}

}
