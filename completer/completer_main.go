package completer

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

const MODULE_MAIN = 0x00
const MODULE_GNB = 0x01
const MODULE_CONFIGURATION = 0x02
const MODULE_SUBSCRIBER = 0x03

var MainSuggestion = []prompt.Suggest{
	{Text: "configuration", Description: "Manage the configuration file"},
	{Text: "subscriber", Description: "Manage the subscribers of the network"},
	{Text: "gnb", Description: "Simulate a 5g gNB"},
	{Text: "exit", Description: "Exit freecli"},
}

// Completer is responsible for the autocompletion of the CLI
func Completer(in prompt.Document) []prompt.Suggest {
	w := in.TextBeforeCursor()
	if strings.HasPrefix(w, "configuration") {
		return completerConfiguration(in)
	}
	if strings.HasPrefix(w, "subscriber") {
		return completerSubscriber(in)
	}
	return prompt.FilterHasPrefix(PromptConfig.Suggestion, w, true)
}
