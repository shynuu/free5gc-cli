package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type PLMNSupportItem struct {
	PLMNIdentity     PLMNIdentity
	SliceSupportList SliceSupportList
	IEExtensions     *ProtocolExtensionContainerPLMNSupportItemExtIEs `aper:"optional"`
}
