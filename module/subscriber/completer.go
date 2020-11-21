package subscriber

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

var SubscriberSuggestion = []prompt.Suggest{
	{Text: "user", Description: "Manage the subscribers of free5gc"},
	{Text: "configuration", Description: "Manage the configuration of the module"},
	{Text: "exit", Description: "Exit the subscriber module"},
}

var userSuggestion = []prompt.Suggest{
	{Text: "random", Description: "Register a number of subscribers with a randomly generated supi for a specific plmn"},
	{Text: "register", Description: "Register a new subscriber with a specific supi and plmn"},
	{Text: "delete", Description: "Delete an exising subscriber"},
	{Text: "list", Description: "List all the registered subscribers"},
	{Text: "flush", Description: "Remove all the registered subscribers from the database"},
}

var configurationSuggestion = []prompt.Suggest{
	{Text: "reload", Description: "Reload the configuration of the subscriber module"},
}

var removeSuggestion = []prompt.Suggest{
	{Text: "supi", Description: "The supi of the subscriber to remove"},
}

var updateSuggestion = []prompt.Suggest{
	{Text: "supi", Description: "The supi of the subscriber to update"},
	{Text: "template", Description: "The template configuration file"},
}

var supiSuggestion = &[]prompt.Suggest{}

var plmnSuggestion = &[]prompt.Suggest{}

func completerUser(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	d := strings.Split(in.TextBeforeCursor(), " ")
	if d[1] == "delete" {
		a = in.GetWordBeforeCursor()
		return prompt.FilterHasPrefix(*supiSuggestion, a, true)
	}
	if d[1] == "register" {
		l := len(d)

		if l >= 7 {
			return []prompt.Suggest{}
		}

		if l >= 6 {
			a = in.GetWordBeforeCursor()
			return prompt.FilterHasPrefix(*plmnSuggestion, a, true)
		}

		if l > 2 && l < 4 {
			a = in.GetWordBeforeCursor()
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--supi", Description: "Specify the supi of the user"},
			}, a, true)
		}

		if l > 4 && l < 6 {
			a = in.GetWordBeforeCursor()
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--plmn", Description: "Specify the plmn id of the network"},
			}, a, true)
		}

		return []prompt.Suggest{}
	}

	if d[1] == "random" {
		l := len(d)

		if l >= 7 {
			return []prompt.Suggest{}
		}

		if l >= 6 {
			a = in.GetWordBeforeCursor()
			return prompt.FilterHasPrefix(*plmnSuggestion, a, true)
		}

		if l > 2 && l < 4 {
			a = in.GetWordBeforeCursor()
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--count", Description: "Specify the number of subscribers to generate"},
			}, a, true)
		}

		if l > 4 {
			a = in.GetWordBeforeCursor()
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--plmn", Description: "Specify the plmn id of the network"},
			}, a, true)
		}

		return []prompt.Suggest{}

	}

	return prompt.FilterHasPrefix(userSuggestion, a, true)
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
