package completer

import "github.com/c-bata/go-prompt"

var PromptConfig Prompt

type Prompt struct {
	Title      string
	Prefix     string
	IsEnable   bool
	Suggestion []prompt.Suggest
	IsModule   bool
	Module     int
}

func Initialize() {

	PromptConfig = Prompt{
		IsEnable:   false,
		IsModule:   false,
		Prefix:     "freecli>>>",
		Suggestion: MainSuggestion,
		Title:      "freecli - a CLI tool to interact and test free5gc",
	}

}
