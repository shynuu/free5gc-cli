package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type CriticalityDiagnosticsIEItem struct {
	IECriticality Criticality
	IEID          ProtocolIEID
	TypeOfError   TypeOfError
	IEExtensions  *ProtocolExtensionContainerCriticalityDiagnosticsIEItemExtIEs `aper:"optional"`
}
