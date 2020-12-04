package api

import "free5gc-cli/logger"

// Initialize the api
func Initialize(path string, force bool) {
	InitConfigFactory(path, force)

	err := NGSetup()
	if err != nil {
		panic("Impossible to join AMF")
	}
	logger.GNBLog.Infoln("gNB Identified on 5G Core Network")
}

// Exit the module
func Exit() {
	amfConn.Close()
	upfConn.Close()
}
