package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type NRCellIdentity struct {
	Value aper.BitString `aper:"sizeLB:36,sizeUB:36"`
}
