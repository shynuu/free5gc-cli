package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type TAIBroadcastEUTRAItem struct {
	TAI                      TAI `aper:"valueExt"`
	CompletedCellsInTAIEUTRA CompletedCellsInTAIEUTRA
	IEExtensions             *ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs `aper:"optional"`
}
