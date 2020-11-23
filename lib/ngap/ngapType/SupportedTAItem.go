package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type SupportedTAItem struct {
	TAC               TAC
	BroadcastPLMNList BroadcastPLMNList
	IEExtensions      *ProtocolExtensionContainerSupportedTAItemExtIEs `aper:"optional"`
}
