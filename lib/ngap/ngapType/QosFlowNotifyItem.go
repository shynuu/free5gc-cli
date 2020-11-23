package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type QosFlowNotifyItem struct {
	QosFlowIdentifier QosFlowIdentifier
	NotificationCause NotificationCause
	IEExtensions      *ProtocolExtensionContainerQosFlowNotifyItemExtIEs `aper:"optional"`
}
