package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type CellIDBroadcastNRItem struct {
	NRCGI        NRCGI                                                  `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCellIDBroadcastNRItemExtIEs `aper:"optional"`
}
