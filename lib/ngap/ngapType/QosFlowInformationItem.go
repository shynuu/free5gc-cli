package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type QosFlowInformationItem struct {
	QosFlowIdentifier QosFlowIdentifier
	DLForwarding      *DLForwarding                                           `aper:"optional"`
	IEExtensions      *ProtocolExtensionContainerQosFlowInformationItemExtIEs `aper:"optional"`
}
