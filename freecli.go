package main

import (
	"fmt"
	"free5gc-cli/freecli"
	"free5gc-cli/u32"
)

// load configuration file.yaml
// load ue file.yaml

func main() {

	var ipv4 = u32.IPV4Header{
		Source:      "216.58.215.35",
		Destination: "216.58.215.35",
		Set: u32.IPV4Fields{
			Destination: true,
		},
	}
	fmt.Println(ipv4.BuildMatches())

	var i = u32.IPV4Header{
		Protocol: 0x11,
		Set: u32.IPV4Fields{
			Protocol: true,
		},
	}
	fmt.Println(i.BuildMatches())

	freecli.Initialize()
	freecli.Run()
}
