package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type PDUSessionResourceModifyItemModCfm struct {
	PDUSessionID                            PDUSessionID
	PDUSessionResourceModifyConfirmTransfer aper.OctetString
	IEExtensions                            *ProtocolExtensionContainerPDUSessionResourceModifyItemModCfmExtIEs `aper:"optional"`
}
