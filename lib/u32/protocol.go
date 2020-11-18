package u32

import (
	"bytes"
	"encoding/binary"
)

// U32 = U32{}
// U32.BuildPacket(IPV4{header: }, TCP{source_port: 80}, UDP{source_port: 80})
// U32.TotalLength
// match = U32.BuildMatch

type Protocol interface {
	GetOffset() Offset
	SetOffset(offset *Offset)
	MoveOffset(offset *Offset)
	BuildMatches() string
	NextHeader() string
}

type Offset struct {
	U32Offset string
	Offset    int
}

func Uint16ToUint8(value uint16) []byte {

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, value)
	if err == nil {
		return buf.Bytes()
	}
	return buf.Bytes()
}

func Uint32ToUint8(value uint32) []byte {

	buf := new(bytes.Buffer)
	err := binary.Write(buf, binary.BigEndian, value)
	if err == nil {
		return buf.Bytes()
	}
	return buf.Bytes()
}
