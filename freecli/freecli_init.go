package freecli

import (
	"free5gc-cli/completer"
	"free5gc-cli/factory"
	"free5gc-cli/lib/MongoDBLibrary"

	"github.com/c-bata/go-prompt"
)

// Initialize freeCli
func Initialize() {

	DefaultCLIConfigPath := "config/freecli.yaml"
	factory.InitConfigFactory(DefaultCLIConfigPath)

	// get config file info from WebUIConfig
	mongodb := factory.FreecliConfig.Configuration.Mongodb

	// Connect to MongoDB
	MongoDBLibrary.SetMongoDB(mongodb.Name, mongodb.Url)

	completer.Initialize()
}

// Run launch a new prompt
func Run() {
	p := prompt.New(
		completer.Executor,
		completer.Completer,
		prompt.OptionTitle("freecli - a simple CLI tool to manage free5gc"),
		prompt.OptionPrefix(completer.PromptConfig.Prefix),
		prompt.OptionLivePrefix(completer.ChangeLivePrefix),
		prompt.OptionPrefixTextColor(prompt.Blue),
		prompt.OptionPreviewSuggestionTextColor(prompt.Blue),
		prompt.OptionSelectedSuggestionBGColor(prompt.LightGray),
		prompt.OptionSuggestionBGColor(prompt.DarkGray))
	p.Run()
}
