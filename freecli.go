package main

import (
	"free5gc-cli/boot"
	"free5gc-cli/freecli"
)

func main() {
	boot.Initialize()
	freecli.Initialize()
	freecli.Run()
}
