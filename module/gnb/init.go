package gnb

import (
	"free5gc-cli/factory"
	"free5gc-cli/lib/MongoDBLibrary"
)

// ue register --supi imsi-20893XXXXXX00
// ue deregister --supi imsi-20893XXXXXX00

// pdu list
// pdu register --supi POPOD --plmn 20893
// pdu
// ===> ipv4, qos profile, sessionid
// pdu release --session <session_id>
// pdu

func Initialize() {
	DefaultCLIConfigPath := "config/freecli.yaml"
	factory.InitConfigFactory(DefaultCLIConfigPath)

	// get config file info from WebUIConfig
	mongodb := factory.FreecliConfig.Configuration.Mongodb

	// Connect to MongoDB
	MongoDBLibrary.SetMongoDB(mongodb.Name, mongodb.Url)
}
