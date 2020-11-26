package freecli

import (
	"free5gc-cli/logger"
	"free5gc-cli/module/gnb"
	"free5gc-cli/module/nf"
	"free5gc-cli/module/qos"
	"free5gc-cli/module/subscriber"
	"os"
	"strings"
)

// Executor parse the CLI and executes the dedicated procedures
func Executor(in string) {

	if PromptConfig.Module == nf.MODULE_NF {
		nf.Executor(in)
	}
	if PromptConfig.Module == gnb.MODULE_GNB {
		gnb.Executor(in)
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
		logger.FreecliLog.Infoln("Loading gNB module...")
		PromptConfig.Suggestion = &gnb.GNBSuggestion
		PromptConfig.IsEnable = true
		PromptConfig.Prefix = "gnb# "
		PromptConfig.IsModule = true
		PromptConfig.Module = gnb.MODULE_GNB
		gnb.Initialize()
		return
	}

	if strings.HasPrefix(in, "qos") {
		logger.FreecliLog.Infoln("Loading QoS module...")
		PromptConfig.Suggestion = &qos.QOSSuggestion
		PromptConfig.IsEnable = true
		PromptConfig.Prefix = "qos# "
		PromptConfig.IsModule = true
		PromptConfig.Module = qos.MODULE_QOS
		qos.Initialize()
		return
	}

	if strings.HasPrefix(in, "nf") {
		logger.FreecliLog.Infoln("Loading Network Function module...")
		PromptConfig.Suggestion = &nf.NFSuggestion
		PromptConfig.IsEnable = true
		PromptConfig.Prefix = "network-function# "
		PromptConfig.IsModule = true
		PromptConfig.Module = nf.MODULE_NF
		nf.Initialize()
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
		logger.FreecliLog.Infoln("Releasing gNB module resources")
		gnb.Exit()
		logger.FreecliLog.Infoln("Releasing subscriber module resources")
		subscriber.Exit()
		logger.FreecliLog.Infoln("Releasing network function module resources")
		nf.Exit()
		logger.FreecliLog.Infoln("Bye Bye !")
		os.Exit(0)
	}

}
