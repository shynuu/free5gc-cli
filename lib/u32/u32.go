package u32

import (
	"fmt"
	"strings"
)

// U32 struct
type U32 struct {
	Protocols []Protocol
	Length    int
	Matches   string
	DSCP      uint8
}

// BuildPacket build Packet Headers
func (u32 *U32) BuildPacket() {
	offset := &Offset{Offset: 0, U32Offset: ""}
	for i := 0; i < len(u32.Protocols); i++ {
		u32.Protocols[i].SetOffset(offset)
		u32.Protocols[i].MoveOffset(offset)
	}
}

// BuildMatches build Matches
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

// BuildCommand generate the IPTABLES command
func (u32 *U32) BuildCommand() string {
	return fmt.Sprintf("sudo iptables -t mangle -A POSTROUTING -m u32 --u32 \"%s\" -j DSCP %d", u32.BuildMatches(), u32.DSCP)
}

func NewU32(protocols *[]Protocol, dscp uint8) *U32 {
	var u32 = U32{Protocols: *protocols, DSCP: dscp}
	u32.BuildPacket()
	return &u32
}
