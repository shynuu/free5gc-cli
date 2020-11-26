package nf

import (
	"strings"

	"github.com/c-bata/go-prompt"
)

var NFSuggestion = []prompt.Suggest{
	{Text: "database", Description: "Manage the database"},
	{Text: "configuration", Description: "Manage the configuration of the NF module"},
	{Text: "exit", Description: "Exit the NF module"},
}

var DatabaseSuggestion = []prompt.Suggest{
	{Text: "drop-collection", Description: "Drop a specific collection"},
	{Text: "flush", Description: "Flush all the data from the database"},
	{Text: "exit", Description: "Exit the NF module"},
}

var CollectionSuggestion = &[]prompt.Suggest{}

func completerDatabase(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	d := strings.Split(in.TextBeforeCursor(), " ")
	if d[1] == "drop-collection" {
		l := len(d)
		if l > 2 && l < 4 {
			return prompt.FilterHasPrefix([]prompt.Suggest{
				{Text: "--collection", Description: "Specify the collection to drop"},
			}, a, true)
		}
		if l > 3 && l < 5 {
			return prompt.FilterHasPrefix(*CollectionSuggestion, a, true)
		}
		if l >= 5 {
			return []prompt.Suggest{}
		}
	}
	if d[1] == "flush" {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix(DatabaseSuggestion, a, true)
}

func completerConfiguration(in prompt.Document) []prompt.Suggest {
	a := in.GetWordBeforeCursor()
	a = strings.TrimSpace(a)
	d := in.TextBeforeCursor()
	if len(strings.Split(d, " ")) > 2 {
		return []prompt.Suggest{}
	}
	return prompt.FilterHasPrefix([]prompt.Suggest{
		{Text: "reload", Description: "Reload the configuration of the subscriber module"},
	}, a, true)
}

func CompleterNF(in prompt.Document) []prompt.Suggest {
	a := in.TextBeforeCursor()
	var split = strings.Split(a, " ")
	w := in.GetWordBeforeCursor()
	if len(split) > 1 {
		var v = split[0]
		if v == "database" {
			return completerDatabase(in)
		}
		if v == "configuration" {
			return completerConfiguration(in)
		}
		return prompt.FilterHasPrefix(NFSuggestion, v, true)
	}
	return prompt.FilterHasPrefix(NFSuggestion, w, true)
}
