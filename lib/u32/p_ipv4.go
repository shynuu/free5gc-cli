package u32

import (
	"encoding/hex"
	"net"
	"strconv"
	"strings"
)

type IPV4Fields struct {
	Version        bool
	IHL            bool
	DSCP           bool
	ECN            bool
	TotalLength    bool
	Identification bool
	Flags          bool
	FragmentOffset bool
	TTL            bool
	Protocol       bool
	HeaderChecksum bool
	Source         bool
	Destination    bool
}

type IPV4Header struct {
	Offset *Offset

	Version        uint8
	IHL            uint8
	DSCP           uint8
	ECN            uint8
	TotalLength    uint16
	Identification uint16
	Flags          uint16
	FragmentOffset uint16
	TTL            uint8
	Protocol       uint8
	HeaderChecksum uint16
	Source         string
	Destination    string
	Set            *IPV4Fields
}

func (ipv4 *IPV4Header) NextHeader() string {
	return strconv.Itoa(ipv4.Offset.Offset) + ">>22&0x3C@"
}

func (ipv4 *IPV4Header) GetOffset() Offset {
	return *ipv4.Offset
}

func (ipv4 *IPV4Header) SetOffset(offset *Offset) {
	ipv4.Offset = &Offset{Offset: offset.Offset, U32Offset: offset.U32Offset}
}

func (ipv4 *IPV4Header) MoveOffset(offset *Offset) {
	offset.Offset = 0
	offset.U32Offset += ipv4.NextHeader()
}

func (ipv4 *IPV4Header) Marshall() []byte {

	var bytes []byte

	bytes = append(bytes, (ipv4.Version<<4)+ipv4.IHL)

	bytes = append(bytes, (ipv4.DSCP<<6)+ipv4.ECN)

	bytes = append(bytes, Uint16ToUint8(ipv4.TotalLength)...)

	bytes = append(bytes, Uint16ToUint8(ipv4.Identification)...)

	var tmp uint16 = (ipv4.Flags << 13) + (ipv4.FragmentOffset >> 3)

	bytes = append(bytes, Uint16ToUint8(tmp)...)

	bytes = append(bytes, ipv4.TTL)

	bytes = append(bytes, ipv4.Protocol)

	bytes = append(bytes, Uint16ToUint8(ipv4.HeaderChecksum)...)

	if ipv4.Source == "" {
		ipv4.Source = "0.0.0.0"
	}

	bytes = append(bytes, net.ParseIP(ipv4.Source).To4()...)

	if ipv4.Destination == "" {
		ipv4.Destination = "0.0.0.0"
	}

	bytes = append(bytes, net.ParseIP(ipv4.Destination).To4()...)

	return bytes

}

func (ipv4 *IPV4Header) BuildMatches() string {

	packet := ipv4.Marshall()
	matches := []string{}

	var i int = 0
	for i < len(packet) {

		match := ""
		mask := "0x"

		for index := 0; index < 4; index++ {

			msk, mtch := ipv4.GetMask(i+index, packet[i+index])
			mask += msk
			match += strings.ToUpper(hex.EncodeToString([]byte{mtch}))
		}

		if mask != "0x00000000" {
			match = ipv4.Offset.U32Offset + strconv.Itoa(ipv4.Offset.Offset+i) + "&" + mask + "=0x" + match
			matches = append(matches, match)
		}

		i += 4

	}

	return strings.Join(matches, " && ")
}

func (ipv4 *IPV4Header) GetMask(offset int, value byte) (string, byte) {
	if offset == 0 {
		if ipv4.Set.Version && ipv4.Set.IHL {
			return "FF", 0xFF & value
		} else if ipv4.Set.Version {
			return "0F", 0x0F & value
		} else if ipv4.Set.IHL {
			return "F0", 0xF0 & value
		}
		return "00", 00

	}
	if offset == 1 {
		if ipv4.Set.ECN && ipv4.Set.DSCP {
			return "FF", 0xFF & value
		} else if ipv4.Set.ECN {
			return "03", 0x0F & value
		} else if ipv4.Set.DSCP {
			return "FC", 0xF0 & value
		}
		return "00", 00

	}
	if offset == 2 || offset == 3 {
		if ipv4.Set.TotalLength {
			return "FF", 0xFF & value
		}
		return "00", 00

	}
	if offset == 4 || offset == 5 {
		if ipv4.Set.Identification {
			return "FF", 0xFF & value
		}
		return "00", 00

	}
	if offset == 6 {
		if ipv4.Set.Flags && ipv4.Set.FragmentOffset {
			return "FF", 0xFF & value
		} else if ipv4.Set.Flags {
			return "E0", 0xE0 & value
		} else if ipv4.Set.FragmentOffset {
			return "1F", 0x1F & value
		}
		return "00", 00
	}
	if offset == 7 {
		if ipv4.Set.FragmentOffset {
			return "FF", 0xFF & value
		}
		return "00", 00
	}
	if offset == 8 {
		if ipv4.Set.TTL {
			return "FF", 0xFF & value
		}
		return "00", 00
	}
	if offset == 9 {
		if ipv4.Set.Protocol {
			return "FF", 0xFF & value
		}
		return "00", 00
	}
	if offset == 10 || offset == 11 {
		if ipv4.Set.TTL {
			return "FF", 0xFF & value
		}
		return "00", 00
	}
	if offset == 12 || offset == 13 || offset == 14 || offset == 15 {
		if ipv4.Set.Source {
			return "FF", 0xFF & value
		}
		return "00", 00
	}
	if offset == 16 || offset == 17 || offset == 18 || offset == 19 {
		if ipv4.Set.Destination {
			return "FF", 0xFF & value
		}
		return "00", 00
	}

	return "", 00
}
