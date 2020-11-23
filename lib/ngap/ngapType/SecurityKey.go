package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type SecurityKey struct {
	Value aper.BitString `aper:"sizeLB:256,sizeUB:256"`
}
