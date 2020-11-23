package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type EPSTAI struct {
	PLMNIdentity PLMNIdentity
	EPSTAC       EPSTAC
	IEExtensions *ProtocolExtensionContainerEPSTAIExtIEs `aper:"optional"`
}
