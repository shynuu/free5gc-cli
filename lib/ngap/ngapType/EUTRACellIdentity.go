package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type EUTRACellIdentity struct {
	Value aper.BitString `aper:"sizeLB:28,sizeUB:28"`
}
