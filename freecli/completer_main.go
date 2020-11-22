package freecli

import (
	"free5gc-cli/module/gnb"
	"free5gc-cli/module/subscriber"

	"github.com/c-bata/go-prompt"
)

const MODULE_MAIN = 0x00

var MainSuggestion = []prompt.Suggest{
	{Text: "subscriber", Description: "Launch the subscriber module"},
	{Text: "gnb", Description: "Launch the gnb emulator module"},
	{Text: "exit", Description: "Exit freecli"},
}

// Completer is responsible for the autocompletion of the CLI
func Completer(in prompt.Document) []prompt.Suggest {
	w := in.TextBeforeCursor()

	if PromptConfig.IsModule && PromptConfig.Module == subscriber.MODULE_SUBSCRIBER {
		return subscriber.CompleterSubscriber(in)
	}

	if PromptConfig.IsModule && PromptConfig.Module == gnb.MODULE_GNB {
		return gnb.CompleterGNB(in)
	}

	return prompt.FilterHasPrefix(*PromptConfig.Suggestion, w, true)
}
