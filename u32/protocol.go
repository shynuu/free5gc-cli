package u32

// U32 = U32{}
// U32.BuildPacket(IPV4{header: }, TCP{source_port: 80}, UDP{source_port: 80})
// U32.TotalLength
// match = U32.BuildMatch

type Protocol interface {
	HeaderLength() int
	GetOffset() int
	SetOffset(start int)
	BuildMatches() string
}

type Field interface {
	Length() int
}
