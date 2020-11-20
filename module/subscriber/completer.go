package subscriber

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

var SubscriberSuggestion = []prompt.Suggest{
	{Text: "user", Description: "Manage the user"},
	{Text: "configuration", Description: "Manage the configuration of the module"},
	{Text: "exit", Description: "Exit the subscriber module"},
}

var userSuggestion = []prompt.Suggest{
	{Text: "register", Description: "Register a new subscriber"},
	{Text: "flush", Description: "Remove all the subscribers from the database"},
	{Text: "refresh", Description: "Refresh the list of registered subscribers in memory"},
	{Text: "remove", Description: "Remove an exising subscriber"},
	{Text: "list", Description: "List all the subscribers"},
	{Text: "update", Description: "Update an exisiting subscribers"},
}

var configurationSuggestion = []prompt.Suggest{
	{Text: "reload", Description: "Reload the configuration of module"},
}

var registerSuggestion = []prompt.Suggest{
	{Text: "reload", Description: "Reload the configuration of module"},
}

var removeSuggestion = []prompt.Suggest{
	{Text: "supi", Description: "The supi of the UE to remove"},
}

var updateSuggestion = []prompt.Suggest{
	{Text: "supi", Description: "The supi of the UE to update"},
	{Text: "template", Description: "The template configuration file"},
}

var supiSuggestion = &[]prompt.Suggest{}

func completerUser(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	d := strings.Split(in.TextBeforeCursor(), " ")
	if d[1] == "remove" {
		a = in.GetWordBeforeCursor()
		return prompt.FilterHasPrefix(*supiSuggestion, a, true)
	}

	return prompt.FilterHasPrefix(userSuggestion, a, true)
}

func completerRegister(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	return prompt.FilterHasPrefix(registerSuggestion, a, true)
}

func completerConfiguration(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	d := in.TextBeforeCursor()
	if len(strings.Split(d, " ")) > 2 {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(configurationSuggestion, a, true)
}

func completerUpdate(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	return prompt.FilterHasPrefix(updateSuggestion, a, true)
}

func completerRemove(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	return prompt.FilterHasPrefix(*supiSuggestion, a, true)
}

func CompleterSubscriber(in prompt.Document) []prompt.Suggest {
	a := in.TextBeforeCursor()
	var split = strings.Split(a, " ")
	w := in.GetWordBeforeCursor()
	if len(split) > 1 {
		var v = split[0]
		if v == "user" {
			return completerUser(in)
		}
		if v == "configuration" {
			return completerConfiguration(in)
		}
		return prompt.FilterHasPrefix(SubscriberSuggestion, v, true)
	}
	return prompt.FilterHasPrefix(SubscriberSuggestion, w, true)
}
