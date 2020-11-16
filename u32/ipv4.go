package u32

import (
	"bytes"
	"encoding/binary"
	"net"
)

type IPV4Header struct {
	Offset uint

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
}

func (ipv4 *IPV4Header) HeaderLength() int {
	return 20
}

func (ipv4 *IPV4Header) GetOffset() uint {
	return ipv4.Offset
}

func (ipv4 *IPV4Header) SetOffset(offset uint) {
	ipv4.Offset = offset
}

func (ipv4 *IPV4Header) BuildMatches(packet []byte) string {

	matches := ""

	return matches
}

func Uint16ToUint8(value uint16) []byte {

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, value)
	if err == nil {
		return buf.Bytes()
	}
	return buf.Bytes()
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

	bytes = append(bytes, net.ParseIP(ipv4.Source).To4()...)

	bytes = append(bytes, net.ParseIP(ipv4.Destination).To4()...)

	return bytes

}
