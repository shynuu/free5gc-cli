package parser

import (
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
)

var mainSuggestion = []prompt.Suggest{
	{Text: "configuration", Description: "Manage the configuration file"},
	{Text: "subscriber", Description: "Manage the subscribers of the network"},
	{Text: "exit", Description: "Exit freecli"},
}

var configurationSuggestion = []prompt.Suggest{
	{Text: "load", Description: "Load a configuration file for every"},
	{Text: "show", Description: "Show the current configuration file"},
}

var subscriberSuggestion = []prompt.Suggest{
	{Text: "register", Description: "Register a new UE"},
	{Text: "remove", Description: "Remove an exising UE"},
	{Text: "update", Description: "Update an exisiting UE"},
}

var registerSuggestion = []prompt.Suggest{
	{Text: "imsi", Description: "The supi of the UE to register"},
}

// Executor parse the CLI and executes the dedicated procedures
func Executor(in string) {

	if in == "exit" {
		os.Exit(0)
	}
}

// Completer is responsible for the autocompletion of the CLI
func Completer(in prompt.Document) []prompt.Suggest {
	w := in.TextBeforeCursor()
	if strings.HasPrefix(w, "configuration") {
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
	if strings.HasPrefix(w, "subscriber") {
		a := in.GetWordBeforeCursorWithSpace()
		a = strings.TrimSpace(a)
		if strings.HasPrefix(a, "register") || strings.HasPrefix(a, "remove") || strings.HasPrefix(a, "update") {
			return []prompt.Suggest{}
		}
		if a == "subscriber" {
			return subscriberSuggestion
		}
		return prompt.FilterHasPrefix(subscriberSuggestion, a, true)
	}
	return prompt.FilterHasPrefix(mainSuggestion, w, true)
}
