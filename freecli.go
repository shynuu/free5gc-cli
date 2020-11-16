package main

import (
	"encoding/hex"
	"fmt"
	"free5gc-cli/freecli"
	"free5gc-cli/u32"
)

// load configuration file.yaml
// load ue file.yaml

func main() {

	var ipv4 = u32.IPV4Header{
		Version:        4,
		IHL:            5,
		TotalLength:    116,
		Identification: 0x77a2,
		Flags:          2,
		FragmentOffset: 0,
		TTL:            64,
		Protocol:       0x11,
		HeaderChecksum: 0x1cdf,
		Source:         "10.200.200.1",
		Destination:    "10.200.200.101",
	}
	fmt.Println(ipv4.Version)
	fmt.Println(hex.EncodeToString(ipv4.Marshall()))

	freecli.Initialize()
	freecli.Run()
}
