package main

import (
	"fmt"
	"free5gc-cli/u32"
)

// load configuration file.yaml
// load ue file.yaml

func main() {

	var U32 = u32.U32{Protocols: []u32.Protocol{
		&u32.IPV4Header{
			Destination: "216.58.215.35",
			Set: u32.IPV4Fields{
				Destination: true,
			},
		},
		&u32.IPV4Header{
			Destination: "216.58.215.36",
			Set: u32.IPV4Fields{
				Destination: true,
			},
		},
		&u32.IPV4Header{
			Destination: "216.58.215.36",
			Set: u32.IPV4Fields{
				Destination: true,
			},
		},
	},
	}

	U32.BuildPacket()
	result := U32.BuildMatches()
	fmt.Println(result)

	// freecli.Initialize()
	// freecli.Run()
}
