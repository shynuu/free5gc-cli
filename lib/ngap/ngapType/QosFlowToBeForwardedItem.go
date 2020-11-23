package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type QosFlowToBeForwardedItem struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowToBeForwardedItemExtIEs `aper:"optional"`
}
