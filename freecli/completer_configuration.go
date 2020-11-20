package freecli

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

var ConfigurationSuggestion = []prompt.Suggest{
	{Text: "load", Description: "Load a default configuration file for "},
	{Text: "show", Description: "Show the current configuration file"},
	{Text: "exit", Description: "Exit the module"},
}

func completerConfiguration(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursorWithSpace()
	a = strings.TrimSpace(a)
	if strings.HasPrefix(a, "load") || strings.HasPrefix(a, "show") {
		return []prompt.Suggest{}
	}
	if a == "configuration" {
		return ConfigurationSuggestion
	}
	return prompt.FilterHasPrefix(ConfigurationSuggestion, a, true)
}
