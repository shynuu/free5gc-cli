package u32

import (
	"strings"
)

type U32 struct {
	Protocols []Protocol
	Length    int
	Matches   string
}

func (u32 *U32) BuildPacket() {
	offset := &Offset{Offset: 0, U32Offset: ""}
	for i := 0; i < len(u32.Protocols); i++ {
		u32.Protocols[i].SetOffset(offset)
		u32.Protocols[i].MoveOffset(offset)
	}
}

func (u32 *U32) BuildMatches() string {
	var matches []string
	for i := 0; i < len(u32.Protocols); i++ {
		match := u32.Protocols[i].BuildMatches()
		if match != "" {
			matches = append(matches, match)
		}
	}
	u32.Matches = strings.Join(matches, " && ")
	return u32.Matches
}

func (u32 *U32) BuildCommand() string {
	return "-m u32 --u32 " + "\"" + u32.BuildMatches() + "\""
}

// var U32 = u32.U32{Protocols: []u32.Protocol{
// 	&u32.IPV4Header{
// 		Source:      "10.200.200.1",
// 		Destination: "10.200.200.102",
// 		Protocol:    u32.PROTO_UDP,
// 		Set: &u32.IPV4Fields{
// 			Source:      true,
// 			Destination: true,
// 			Protocol:    true,
// 		},
// 	},
// 	&u32.UDPHeader{
// 		SourcePort:      2152,
// 		DestinationPort: 2152,
// 		Set: &u32.UDPFields{
// 			SourcePort:      true,
// 			DestinationPort: true,
// 		},
// 	},
// 	&u32.GTPv1Header{
// 		TEID: 1,
// 		Set:  &u32.GTPv1Fields{TEID: true},
// 	},
// },
// }

// U32.BuildPacket()
// result := U32.BuildCommand()
// fmt.Println("sudo iptables -t mangle -A POSTROUTING " + result + " -j DSCP --set-dscp 63")
