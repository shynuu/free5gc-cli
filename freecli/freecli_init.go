package freecli

import (
	"github.com/c-bata/go-prompt"
)

// Initialize freeCli
func Initialize() {

	InitializePrompt()
}

// Run launch a new prompt
func Run() {
	p := prompt.New(
		Executor,
		Completer,
		prompt.OptionTitle(PromptConfig.Title),
		prompt.OptionPrefix(PromptConfig.Prefix),
		prompt.OptionLivePrefix(ChangeLivePrefix),
		prompt.OptionPrefixTextColor(prompt.Blue),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray))
	p.Run()
}
