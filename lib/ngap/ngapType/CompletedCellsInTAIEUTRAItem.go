package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type CompletedCellsInTAIEUTRAItem struct {
	EUTRACGI     EUTRACGI                                                      `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCompletedCellsInTAIEUTRAItemExtIEs `aper:"optional"`
}
