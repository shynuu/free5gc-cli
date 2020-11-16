package u32

import (
	"strings"
)

type U32 struct {
	Protocols []Protocol
	Length    int
	Matches   string
}

func (u32 *U32) BuildPacket(protocols ...Protocol) {
	offset := 0
	for i := 0; i < len(protocols); i++ {
		protocols[i].SetOffset(offset)
		offset += protocols[i].HeaderLength()
	}
}

func (u32 *U32) BuildMatches(protocols ...Protocol) string {
	var matches []string
	for i := 0; i < len(protocols); i++ {
		match := protocols[i].BuildMatches()
		if match != "" {
			matches = append(matches, match)
		}
	}
	u32.Matches = strings.Join(matches, " & ")
	return u32.Matches
}
