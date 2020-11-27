package main

import (
	"free5gc-cli/freecli"
)

// load configuration file.yaml
// load ue file.yaml

func main() {

	freecli.Initialize()
	freecli.Run()

}
