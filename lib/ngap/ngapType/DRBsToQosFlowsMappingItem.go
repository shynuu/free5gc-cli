package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type DRBsToQosFlowsMappingItem struct {
	DRBID                 DRBID
	AssociatedQosFlowList AssociatedQosFlowList
	IEExtensions          *ProtocolExtensionContainerDRBsToQosFlowsMappingItemExtIEs `aper:"optional"`
}
