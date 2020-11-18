package u32

import (
	"encoding/hex"
	"strconv"
	"strings"
)

type TCPFields struct {
	SourcePort      bool
	DestinationPort bool
	SequenceNumber  bool
	ACKNumber       bool
	DataOffset      bool
	Flags           bool
	WindowSize      bool
	Checksum        bool
	UrgentPointer   bool
}

type TCPHeader struct {
	Offset *Offset

	SourcePort      uint16
	DestinationPort uint16
	SequenceNumber  uint32
	ACKNumber       uint32
	DataOffset      uint8
	Flags           uint16
	WindowSize      uint16
	Checksum        uint16
	UrgentPointer   uint16

	Set *TCPFields
}

func (tcp *TCPHeader) NextHeader() string {
	return strconv.Itoa(tcp.Offset.Offset+12) + ">>26&0x3C@"
}

func (tcp *TCPHeader) GetOffset() Offset {
	return *tcp.Offset
}

func (tcp *TCPHeader) SetOffset(offset *Offset) {
	tcp.Offset = &Offset{Offset: offset.Offset, U32Offset: offset.U32Offset}
}

func (tcp *TCPHeader) MoveOffset(offset *Offset) {
	offset.Offset = 0
	offset.U32Offset += tcp.NextHeader()
}

func (tcp *TCPHeader) Marshall() []byte {

	var bytes []byte

	bytes = append(bytes, Uint16ToUint8(tcp.SourcePort)...)

	bytes = append(bytes, Uint16ToUint8(tcp.DestinationPort)...)

	bytes = append(bytes, Uint32ToUint8(tcp.SequenceNumber)...)

	bytes = append(bytes, Uint32ToUint8(tcp.ACKNumber)...)

	var tmp = Uint16ToUint8(tcp.Flags)

	bytes = append(bytes, (tcp.DataOffset<<4)+tmp[0])

	bytes = append(bytes, tmp[1])

	bytes = append(bytes, Uint16ToUint8(tcp.WindowSize)...)

	bytes = append(bytes, Uint16ToUint8(tcp.Checksum)...)

	bytes = append(bytes, Uint16ToUint8(tcp.UrgentPointer)...)

	bytes = append(bytes, []byte{0x0, 0x0, 0x0, 0x0}...)

	return bytes

}

func (tcp *TCPHeader) BuildMatches() string {

	packet := tcp.Marshall()
	matches := []string{}

	var i int = 0
	for i < len(packet) {

		match := ""
		mask := "0x"

		for index := 0; index < 4; index++ {

			msk, mtch := tcp.GetMask(i+index, packet[i+index])
			mask += msk
			match += strings.ToUpper(hex.EncodeToString([]byte{mtch}))
		}

		if mask != "0x00000000" {
			match = tcp.Offset.U32Offset + strconv.Itoa(tcp.Offset.Offset+i) + "&" + mask + "=0x" + match
			matches = append(matches, match)
		}

		i += 4

	}

	return strings.Join(matches, " && ")
}

func (tcp *TCPHeader) GetMask(offset int, value byte) (string, byte) {
	if offset == 0 || offset == 1 {
		if tcp.Set.SourcePort {
			return "FF", 0xFF & value
		}
		return "00", 00
	}

	if offset == 2 || offset == 3 {
		if tcp.Set.DestinationPort {
			return "FF", 0xFF & value
		}
		return "00", 00
	}

	if offset == 4 || offset == 5 || offset == 6 || offset == 7 {
		if tcp.Set.SequenceNumber {
			return "FF", 0xFF & value
		}
		return "00", 00
	}

	if offset == 8 || offset == 9 || offset == 10 || offset == 11 {
		if tcp.Set.ACKNumber {
			return "FF", 0xFF & value
		}
		return "00", 00
	}

	if offset == 12 {
		if tcp.Set.DataOffset && tcp.Set.Flags {
			return "F1", 0xF1 & value
		} else if tcp.Set.DataOffset {
			return "F0", 0xF0 & value
		} else if tcp.Set.Flags {
			return "0F", 0x0F & value
		}
		return "00", 00
	}

	if offset == 13 {
		if tcp.Set.Flags {
			return "FF", 0xFF & value
		}
		return "00", 00
	}

	if offset == 14 || offset == 15 {
		if tcp.Set.WindowSize {
			return "FF", 0xFF & value
		}
		return "00", 00
	}

	if offset == 16 || offset == 17 {
		if tcp.Set.Checksum {
			return "FF", 0xFF & value
		}
		return "00", 00
	}

	if offset == 18 || offset == 19 {
		if tcp.Set.UrgentPointer {
			return "FF", 0xFF & value
		}
		return "00", 00
	}

	return "00", 00
}
