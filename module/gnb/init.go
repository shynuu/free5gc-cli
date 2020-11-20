package gnb

import (
	"free5gc-cli/factory"
	"free5gc-cli/lib/MongoDBLibrary"
)

func Initialize() {
	DefaultCLIConfigPath := "config/freecli.yaml"
	factory.InitConfigFactory(DefaultCLIConfigPath)

	// get config file info from WebUIConfig
	mongodb := factory.FreecliConfig.Configuration.Mongodb

	// Connect to MongoDB
	MongoDBLibrary.SetMongoDB(mongodb.Name, mongodb.Url)
}
