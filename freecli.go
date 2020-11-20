package main

import (
	"fmt"
	"free5gc-cli/freecli"
	"free5gc-cli/lib/u32"
)

// load configuration file.yaml
// load ue file.yaml

func main() {

	var U32 = u32.U32{Protocols: []u32.Protocol{
		&u32.IPV4Header{
			Source:      "10.200.200.1",
			Destination: "10.200.200.102",
			Protocol:    u32.PROTO_UDP,
			Set: &u32.IPV4Fields{
				Source:      true,
				Destination: true,
				Protocol:    true,
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
	},
	}

	U32.BuildPacket()
	result := U32.BuildCommand()
	fmt.Println("sudo iptables -t mangle -A POSTROUTING " + result + " -j DSCP --set-dscp 63")

	freecli.Initialize()
	freecli.Run()
}
