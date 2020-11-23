package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type QosFlowSetupResponseItemSURes struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowSetupResponseItemSUResExtIEs `aper:"optional"`
}
