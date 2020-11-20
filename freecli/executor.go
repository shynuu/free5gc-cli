package freecli

import (
	"free5gc-cli/logger"
	"free5gc-cli/module/gnb"
	"free5gc-cli/module/subscriber"
	"os"
	"strings"
)

// Executor parse the CLI and executes the dedicated procedures
func Executor(in string) {

	if PromptConfig.Module == gnb.MODULE_GNB {

	}

	if PromptConfig.Module == subscriber.MODULE_SUBSCRIBER {
		subscriber.Executor(in)
	}

	if strings.HasPrefix(in, "subscriber") {
		logger.FreecliLog.Infoln("Loading subscriber module...")
		PromptConfig.Suggestion = &subscriber.SubscriberSuggestion
		PromptConfig.IsEnable = true
		PromptConfig.Prefix = "subscriber# "
		PromptConfig.IsModule = true
		PromptConfig.Module = subscriber.MODULE_SUBSCRIBER
		subscriber.Initialize()
		return
	}

	if strings.HasPrefix(in, "gnb") {
		logger.FreecliLog.Infoln("Loading g-nb module...")
		PromptConfig.Suggestion = &gnb.GNBSuggestion
		PromptConfig.IsEnable = true
		PromptConfig.Prefix = "gnb# "
		PromptConfig.IsModule = true
		PromptConfig.Module = gnb.MODULE_GNB
		return
	}

	if in == "exit" {
		if PromptConfig.IsModule {
			PromptConfig.Suggestion = &MainSuggestion
			PromptConfig.IsEnable = true
			PromptConfig.Prefix = "freecli> "
			PromptConfig.IsModule = false
			PromptConfig.Module = MODULE_MAIN
			logger.FreecliLog.Infoln("Exiting Module...")
			return
		}
		logger.FreecliLog.Infoln("Exiting Freecli...")
		os.Exit(0)
	}

}
