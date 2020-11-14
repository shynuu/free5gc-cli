package completer

import (
	"fmt"
	"free5gc-cli/logger"
	"os"
	"strings"
)

// Executor parse the CLI and executes the dedicated procedures
func Executor(in string) {

	if in == "subscriber" {
		logger.FreecliLog.Infoln("Loading subscriber module...")
		PromptConfig.Suggestion = SubscriberSuggestion
		PromptConfig.IsEnable = true
		PromptConfig.Prefix = "subscriber>"
		PromptConfig.IsModule = true
		PromptConfig.Module = MODULE_SUBSCRIBER
		return
	}

	if in == "configuration" {
		logger.FreecliLog.Infoln("Loading configuration module...")
		PromptConfig.Suggestion = ConfigurationSuggestion
		PromptConfig.IsEnable = true
		PromptConfig.Prefix = "configuration>"
		PromptConfig.IsModule = true
		PromptConfig.Module = MODULE_CONFIGURATION
		return
	}

	if in == "gnb" {
		logger.FreecliLog.Infoln("Loading g-nb module...")
		PromptConfig.Suggestion = GNBSuggestion
		PromptConfig.IsEnable = true
		PromptConfig.Prefix = "gnb>"
		PromptConfig.IsModule = true
		PromptConfig.Module = MODULE_CONFIGURATION
		return
	}

	if in == "exit" {
		if PromptConfig.IsModule {
			PromptConfig.Suggestion = MainSuggestion
			PromptConfig.IsEnable = true
			PromptConfig.Prefix = "free5gc>>>"
			PromptConfig.IsModule = false
			PromptConfig.Module = MODULE_MAIN
			logger.FreecliLog.Infoln("Exiting Module...")
			return
		}
		logger.FreecliLog.Infoln("Exiting Freecli...")
		os.Exit(0)
	}

	if strings.HasPrefix(in, "subscriber") {
		fmt.Println("subscriber")
	}
}
