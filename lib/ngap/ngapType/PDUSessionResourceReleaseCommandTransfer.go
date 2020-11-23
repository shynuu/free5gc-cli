package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type PDUSessionResourceReleaseCommandTransfer struct {
	Cause        Cause                                                                     `aper:"valueLB:0,valueUB:5"`
	IEExtensions *ProtocolExtensionContainerPDUSessionResourceReleaseCommandTransferExtIEs `aper:"optional"`
}
