package freecli

import (
	"free5gc-cli/factory"
	"free5gc-cli/lib/MongoDBLibrary"
	"free5gc-cli/parser"

	"github.com/c-bata/go-prompt"
)

// Initialize freeCli
func Initialize() {

	DefaultWebUIConfigPath := "config/freecli.yaml"
	factory.InitConfigFactory(DefaultWebUIConfigPath)

	// get config file info from WebUIConfig
	mongodb := factory.FreecliConfig.Configuration.Mongodb

	// Connect to MongoDB
	MongoDBLibrary.SetMongoDB(mongodb.Name, mongodb.Url)
}

func Run() {
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
