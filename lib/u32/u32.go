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
