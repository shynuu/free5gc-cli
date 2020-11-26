package nf

import (
	"free5gc-cli/logger"
	"strings"
)

func executorDatabase(in string) {
	cmd := strings.Split(strings.TrimSpace(in), " ")
	l := len(cmd)

	if l < 2 {
		return
	}

	if cmd[1] == "drop-collection" && l == 4 {
		logger.NFLog.Infoln("Dropping collection", cmd[3])
		nf.DropDatabase(cmd[3])
	}

	if cmd[1] == "flush" && l == 2 {
		logger.NFLog.Infoln("Dropping complete database")
		nf.FlushDatabase()
	}
}

func executorConfiguration(in string) {
	s := strings.TrimSpace(in)
	if s == "configuration reload" {
		Reload()
	}
}

// Executor parse CLI
func Executor(in string) {

	if strings.HasPrefix(in, "configuration") {
		executorConfiguration(in)
	}

	if strings.HasPrefix(in, "database") {
		executorDatabase(in)
	}

}
