package main

import (
	"free5gc-cli/parser"

	"github.com/c-bata/go-prompt"
)

// load configuration file.yaml
// load ue file.yaml

func main() {
	p := prompt.New(
		parser.Executor,
		parser.Completer,
		prompt.OptionTitle("freecli - a simple CLI to manage free5gc"),
		prompt.OptionPrefix("freecli>"),
		prompt.OptionPrefixTextColor(prompt.Blue),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray))
	p.Run()
}
