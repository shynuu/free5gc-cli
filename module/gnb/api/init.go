package api

func Initialize(path string, force bool) {
	InitConfigFactory(path, force)
}

func Exit() {
	amfConn.Close()
	upfConn.Close()
}
