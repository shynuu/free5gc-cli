package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type RANStatusTransferTransparentContainer struct {
	DRBsSubjectToStatusTransferList DRBsSubjectToStatusTransferList
	IEExtensions                    *ProtocolExtensionContainerRANStatusTransferTransparentContainerExtIEs `aper:"optional"`
}
