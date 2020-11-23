package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type PDUSessionResourceModifyItemModInd struct {
	PDUSessionID                               PDUSessionID
	PDUSessionResourceModifyIndicationTransfer aper.OctetString
	IEExtensions                               *ProtocolExtensionContainerPDUSessionResourceModifyItemModIndExtIEs `aper:"optional"`
}
