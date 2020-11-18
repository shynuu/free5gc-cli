package u32

import (
	"encoding/hex"
	"strconv"
	"strings"
)

type GTPv1Fields struct {
	Version        bool
	ProtocolType   bool
	Flags          bool
	MessageType    bool
	Length         bool
	TEID           bool
	SequenceNumber bool
	NPDU           bool
	NextHeaderType bool
}

type GTPv1Header struct {
	Offset *Offset

	Version        uint8
	ProtocolType   uint8
	Flags          uint8
	MessageType    uint8
	Length         uint16
	TEID           uint32
	SequenceNumber uint16
	NPDU           uint8
	NextHeaderType uint8

	Set *GTPv1Fields
}

func (gtp *GTPv1Header) NextHeader() string {
	return ""
}

func (gtp *GTPv1Header) GetOffset() Offset {
	return *gtp.Offset
}

func (gtp *GTPv1Header) SetOffset(offset *Offset) {
	gtp.Offset = &Offset{Offset: offset.Offset, U32Offset: offset.U32Offset}
}

func (gtp *GTPv1Header) MoveOffset(offset *Offset) {
	offset.Offset += 12
}

func (gtp *GTPv1Header) Marshall() []byte {

	gtp.Version = 1
	gtp.ProtocolType = 1
	gtp.Set.Version = true
	gtp.Set.ProtocolType = true

	var bytes []byte

	bytes = append(bytes, ((gtp.Version<<5)+(gtp.ProtocolType<<4))+gtp.Flags)

	bytes = append(bytes, gtp.MessageType)

	bytes = append(bytes, Uint16ToUint8(gtp.Length)...)

	bytes = append(bytes, Uint32ToUint8(gtp.TEID)...)

	bytes = append(bytes, Uint16ToUint8(gtp.SequenceNumber)...)

	bytes = append(bytes, gtp.NPDU)

	bytes = append(bytes, gtp.NextHeaderType)

	return bytes

}

func (gtp *GTPv1Header) BuildMatches() string {

	packet := gtp.Marshall()
	matches := []string{}

	var i int = 0
	for i < len(packet) {

		match := ""
		mask := "0x"

		for index := 0; index < 4; index++ {

			msk, mtch := gtp.GetMask(i+index, packet[i+index])
			mask += msk
			match += strings.ToUpper(hex.EncodeToString([]byte{mtch}))
		}

		if mask != "0x00000000" {
			match = gtp.Offset.U32Offset + strconv.Itoa(gtp.Offset.Offset+i) + "&" + mask + "=0x" + match
			matches = append(matches, match)
		}

		i += 4

	}

	return strings.Join(matches, " && ")
}

func (gtp *GTPv1Header) GetMask(offset int, value byte) (string, byte) {
	if offset == 0 {
		if gtp.Set.Version && gtp.Set.ProtocolType && gtp.Set.Flags {
			return "FF", 0xFF & value
		} else if gtp.Set.Version && gtp.Set.ProtocolType {
			return "F0", 0xF0 & value
		}
		return "00", 00

	}
	if offset == 1 {
		if gtp.Set.MessageType {
			return "FF", 0xFF & value
		}
		return "00", 00

	}
	if offset == 2 || offset == 3 {
		if gtp.Set.Length {
			return "FF", 0xFF & value
		}
		return "00", 00
	}
	if offset == 4 || offset == 5 || offset == 6 || offset == 7 {
		if gtp.Set.TEID {
			return "FF", 0xFF & value
		}
		return "00", 00
	}
	if offset == 8 || offset == 9 {
		if gtp.Set.SequenceNumber {
			return "FF", 0xFF & value
		}
		return "00", 00
	}
	if offset == 10 {
		if gtp.Set.NPDU {
			return "FF", 0xFF & value
		}
		return "00", 00
	}
	if offset == 11 {
		if gtp.Set.NextHeaderType {
			return "FF", 0xFF & value
		}
		return "00", 00
	}
	return "", 00
}
