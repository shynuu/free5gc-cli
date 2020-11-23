package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type AssociatedQosFlowItem struct {
	QosFlowIdentifier        QosFlowIdentifier
	QosFlowMappingIndication *aper.Enumerated                                       `aper:"optional"`
	IEExtensions             *ProtocolExtensionContainerAssociatedQosFlowItemExtIEs `aper:"optional"`
}
