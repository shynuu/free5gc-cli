package freecli

import "github.com/c-bata/go-prompt"

var PromptConfig *Prompt

type Prompt struct {
	Title      string
	Prefix     string
	IsEnable   bool
	Suggestion *[]prompt.Suggest
	IsModule   bool
	Module     int
}

func InitializePrompt() {

	PromptConfig = &Prompt{
		IsEnable:   false,
		IsModule:   false,
		Prefix:     "freecli> ",
		Suggestion: &MainSuggestion,
		Title:      "freecli - A CLI tool for free5gc",
	}

}
