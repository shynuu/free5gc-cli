package freecli

import (
	"free5gc-cli/module/gnb"
	"free5gc-cli/module/qos"
	"free5gc-cli/module/subscriber"

	"github.com/c-bata/go-prompt"
)

const MODULE_MAIN = 0x00

var MainSuggestion = []prompt.Suggest{
	{Text: "subscriber", Description: "Load the subscriber module"},
	{Text: "gnb", Description: "Load the gNB emulator module"},
	{Text: "qos", Description: "Load the QoS  module"},
	{Text: "exit", Description: "Exit freecli"},
}

// Completer is responsible for the autocompletion of the CLI
func Completer(in prompt.Document) []prompt.Suggest {
	if PromptConfig.IsModule && PromptConfig.Module == subscriber.MODULE_SUBSCRIBER {
		return subscriber.CompleterSubscriber(in)
	}

	if PromptConfig.IsModule && PromptConfig.Module == gnb.MODULE_GNB {
		return gnb.CompleterGNB(in)
	}

	if PromptConfig.IsModule && PromptConfig.Module == qos.MODULE_QOS {
		return qos.CompleterQOS(in)
	}

	w := in.TextBeforeCursor()
	return prompt.FilterHasPrefix(*PromptConfig.Suggestion, w, true)
}
