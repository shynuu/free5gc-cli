package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type NRCGI struct {
	PLMNIdentity   PLMNIdentity
	NRCellIdentity NRCellIdentity
	IEExtensions   *ProtocolExtensionContainerNRCGIExtIEs `aper:"optional"`
}
