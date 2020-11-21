package subscriber

import (
	"free5gc-cli/lib/MongoDBLibrary"

	"github.com/c-bata/go-prompt"
)

// Initialize the module
func Initialize() {
	DefaultCLIConfigPath := "config/" + MODULE_NAME + ".yaml"
	InitConfigFactory(DefaultCLIConfigPath, false)

	// get config file info from WebUIConfig
	mongodb := SubscriberConfig.Configuration.Mongodb

	// Connect to MongoDB
	MongoDBLibrary.SetMongoDB(mongodb.Name, mongodb.Url)

	DefaultUEConfigPath := "config/" + MODULE_UE_NAME + ".yaml"
	InitializeUEConfiguration(DefaultUEConfigPath, false)

	var l []prompt.Suggest
	for _, plmn := range SubscriberConfig.PLMN.Plmn {
		l = append(l, prompt.Suggest{Text: plmn, Description: ""})
	}
	plmnSuggestion = &l

}

// Reload the module
func Reload() {
	DefaultCLIConfigPath := "config/" + MODULE_NAME + ".yaml"
	InitConfigFactory(DefaultCLIConfigPath, true)

	// get config file info from WebUIConfig
	mongodb := SubscriberConfig.Configuration.Mongodb

	// Connect to MongoDB
	MongoDBLibrary.SetMongoDB(mongodb.Name, mongodb.Url)

	DefaultUEConfigPath := "config/" + MODULE_UE_NAME + ".yaml"
	InitializeUEConfiguration(DefaultUEConfigPath, true)

}
