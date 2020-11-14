package parser

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

var subscriberSuggestion = []prompt.Suggest{
	{Text: "add", Description: "Add a new subscriber"},
	{Text: "refresh", Description: "Refresh the list of registered UE in memory"},
	{Text: "remove", Description: "Remove an exising UE"},
	{Text: "update", Description: "Update an exisiting UE"},
}

var registerSuggestion = []prompt.Suggest{
	{Text: "template", Description: "The template configuration file"},
}

var removeSuggestion = []prompt.Suggest{
	{Text: "supi", Description: "The supi of the UE to remove"},
}

var updateSuggestion = []prompt.Suggest{
	{Text: "supi", Description: "The supi of the UE to update"},
	{Text: "template", Description: "The template configuration file"},
}

var supiSuggestion = []prompt.Suggest{
	{Text: "imsi-2089300007487"},
	{Text: "imsi-2089300007488"},
	{Text: "imsi-2089300007489"},
}

func completerRegister(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	return prompt.FilterHasPrefix(registerSuggestion, a, true)
}

func completerUpdate(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	return prompt.FilterHasPrefix(updateSuggestion, a, true)
}

func completerRemove(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	return prompt.FilterHasPrefix(supiSuggestion, a, true)
}

func completerSubscriber(in prompt.Document) []prompt.Suggest {
	a := in.TextBeforeCursor()
	var split = strings.Split(a, " ")
	w := in.GetWordBeforeCursor()
	if len(split) > 2 {
		var v = split[1]
		if v == "add" {
			return completerRegister(in)
		}
		if v == "remove" {
			return completerRemove(in)
		}
		if v == "update" {
			return completerUpdate(in)
		}
		if v == "refresh" {
			return []prompt.Suggest{}
		}
		return prompt.FilterHasPrefix(subscriberSuggestion, v, true)
	}
	return prompt.FilterHasPrefix(subscriberSuggestion, w, true)
}
