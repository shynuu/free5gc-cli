package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type MultipleTNLInformation struct {
	TNLInformationList TNLInformationList
	IEExtensions       *ProtocolExtensionContainerMultipleTNLInformationExtIEs `aper:"optional"`
}
