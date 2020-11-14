package parser

import (
	"free5gc-cli/logger"
	"os"
)

// Executor parse the CLI and executes the dedicated procedures
func Executor(in string) {

	if in == "exit" {
		logger.AppLog.Infoln("Exiting Freecli...")
		os.Exit(0)
	}
}
