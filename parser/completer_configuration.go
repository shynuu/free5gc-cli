package parser

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

var configurationSuggestion = []prompt.Suggest{
	{Text: "load", Description: "Load a configuration file for every"},
	{Text: "show", Description: "Show the current configuration file"},
}

func completerConfiguration(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursorWithSpace()
	a = strings.TrimSpace(a)
	if strings.HasPrefix(a, "load") || strings.HasPrefix(a, "show") {
		return []prompt.Suggest{}
	}
	if a == "configuration" {
		return configurationSuggestion
	}
	return prompt.FilterHasPrefix(configurationSuggestion, a, true)
}
