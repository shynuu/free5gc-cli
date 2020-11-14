package parser

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

var mainSuggestion = []prompt.Suggest{
	{Text: "configuration", Description: "Manage the configuration file"},
	{Text: "subscriber", Description: "Manage the subscribers of the network"},
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
	return prompt.FilterHasPrefix(mainSuggestion, w, true)
}
