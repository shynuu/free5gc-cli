package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type PDUSessionResourceSetupItemCxtRes struct {
	PDUSessionID                            PDUSessionID
	PDUSessionResourceSetupResponseTransfer aper.OctetString
	IEExtensions                            *ProtocolExtensionContainerPDUSessionResourceSetupItemCxtResExtIEs `aper:"optional"`
}
