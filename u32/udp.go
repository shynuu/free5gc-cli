package u32

type UDPFields struct {
	SourcePort      bool
	DestinationPort bool
	Length          bool
	Checksum        bool
}

type UDPHeader struct {
	Offset int

	SourcePort      uint16
	DestinationPort uint16
	Length          uint16
	Checksum        uint16
}
