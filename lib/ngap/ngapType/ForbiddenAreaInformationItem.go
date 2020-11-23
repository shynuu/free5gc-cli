package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type ForbiddenAreaInformationItem struct {
	PLMNIdentity  PLMNIdentity
	ForbiddenTACs ForbiddenTACs
	IEExtensions  *ProtocolExtensionContainerForbiddenAreaInformationItemExtIEs `aper:"optional"`
}
