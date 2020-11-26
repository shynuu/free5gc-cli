package nf

import (
	"free5gc-cli/lib/MongoDBLibrary"
	"free5gc-cli/module/nf/api"

	"github.com/c-bata/go-prompt"
)

func Initialize() {
	DefaultCLIConfigPath := "config/" + MODULE_NAME + ".yaml"
	InitConfigFactory(DefaultCLIConfigPath, false)

	// get config file info from WebUIConfig
	mongodb := NFConfig.Configuration.Mongodb

	// Connect to MongoDB
	MongoDBLibrary.SetMongoDB(mongodb.Name, mongodb.Url)

	var l []prompt.Suggest
	for _, collection := range api.DatabaseCollectionList {
		l = append(l, prompt.Suggest{Text: collection, Description: ""})
	}
	CollectionSuggestion = &l

	nf = &NF{}

}

func Reload() {
	DefaultCLIConfigPath := "config/" + MODULE_NAME + ".yaml"
	InitConfigFactory(DefaultCLIConfigPath, true)

	// get config file info from WebUIConfig
	mongodb := NFConfig.Configuration.Mongodb

	// Connect to MongoDB
	MongoDBLibrary.SetMongoDB(mongodb.Name, mongodb.Url)

	var l []prompt.Suggest
	for _, collection := range api.DatabaseCollectionList {
		l = append(l, prompt.Suggest{Text: collection, Description: ""})
	}
	CollectionSuggestion = &l

	nf = &NF{}

}

func Exit() {

}
