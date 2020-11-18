package main

import (
	"fmt"
	"free5gc-cli/lib/u32"
)

// load configuration file.yaml
// load ue file.yaml

func main() {

	var U32 = u32.U32{Protocols: []u32.Protocol{
		&u32.IPV4Header{
			Source:      "10.200.200.1",
			Destination: "10.200.200.102",
			Set: &u32.IPV4Fields{
				Source:      true,
				Destination: true,
			},
		},
		&u32.UDPHeader{
			SourcePort:      2152,
			DestinationPort: 2152,
			Set: &u32.UDPFields{
				SourcePort:      true,
				DestinationPort: true,
			},
		},
		&u32.GTPv1Header{
			TEID: 1,
			Set:  &u32.GTPv1Fields{TEID: true},
		},
		&u32.IPV4Header{
			Protocol: u32.PROTO_TCP,
			Set: &u32.IPV4Fields{
				Protocol: true,
			},
		},
		&u32.TCPHeader{
			DestinationPort: 80,
			Flags:           0b100000001,
			Set: &u32.TCPFields{
				DestinationPort: true,
				Flags:           true,
			},
		},
	},
	}

	U32.BuildPacket()
	result := U32.BuildCommand()
	fmt.Println(result)

	// freecli.Initialize()
	// freecli.Run()
}
