package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type FiveGTMSI struct {
	Value aper.OctetString `aper:"sizeLB:4,sizeUB:4"`
}
