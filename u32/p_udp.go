package u32

import (
	"encoding/hex"
	"strconv"
	"strings"
)

type UDPFields struct {
	SourcePort      bool
	DestinationPort bool
	Length          bool
	Checksum        bool
}

type UDPHeader struct {
	Offset *Offset

	SourcePort      uint16
	DestinationPort uint16
	Length          uint16
	Checksum        uint16

	Set *UDPFields
}

func (udp *UDPHeader) NextHeader() string {
	return ""
}

func (udp *UDPHeader) GetOffset() Offset {
	return *udp.Offset
}

func (udp *UDPHeader) SetOffset(offset *Offset) {
	udp.Offset = &Offset{Offset: offset.Offset, U32Offset: offset.U32Offset}
}

func (udp *UDPHeader) MoveOffset(offset *Offset) {
	offset.Offset += 8
	offset.U32Offset += udp.NextHeader()
}

func (udp *UDPHeader) Marshall() []byte {

	var bytes []byte

	bytes = append(bytes, Uint16ToUint8(udp.SourcePort)...)

	bytes = append(bytes, Uint16ToUint8(udp.DestinationPort)...)

	bytes = append(bytes, Uint16ToUint8(udp.Length)...)

	bytes = append(bytes, Uint16ToUint8(udp.Checksum)...)

	return bytes

}

func (udp *UDPHeader) BuildMatches() string {

	packet := udp.Marshall()
	matches := []string{}

	var i int = 0
	for i < len(packet) {

		match := ""
		mask := "0x"

		for index := 0; index < 4; index++ {

			msk, mtch := udp.GetMask(i+index, packet[i+index])
			mask += msk
			match += strings.ToUpper(hex.EncodeToString([]byte{mtch}))
		}

		if mask != "0x00000000" {
			match = udp.Offset.U32Offset + strconv.Itoa(udp.Offset.Offset+i) + "&" + mask + "=0x" + match
			matches = append(matches, match)
		}

		i += 4

	}

	return strings.Join(matches, " && ")
}

func (udp *UDPHeader) GetMask(offset int, value byte) (string, byte) {
	if offset == 0 || offset == 1 {
		if udp.Set.SourcePort {
			return "FF", 0xFF & value
		}
		return "00", 00

	}
	if offset == 2 || offset == 3 {
		if udp.Set.DestinationPort {
			return "FF", 0xFF & value
		}
		return "00", 00

	}
	if offset == 4 || offset == 5 {
		if udp.Set.Length {
			return "FF", 0xFF & value
		}
		return "00", 00

	}
	if offset == 6 || offset == 7 {
		if udp.Set.Checksum {
			return "FF", 0xFF & value
		}
		return "00", 00

	}
	return "", 00
}
