package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type EmergencyAreaIDBroadcastNRItem struct {
	EmergencyAreaID       EmergencyAreaID
	CompletedCellsInEAINR CompletedCellsInEAINR
	IEExtensions          *ProtocolExtensionContainerEmergencyAreaIDBroadcastNRItemExtIEs `aper:"optional"`
}
