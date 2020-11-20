package subscriber

import (
	"free5gc-cli/lib/MongoDBLibrary"
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
