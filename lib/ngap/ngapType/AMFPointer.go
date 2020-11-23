package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type AMFPointer struct {
	Value aper.BitString `aper:"sizeLB:6,sizeUB:6"`
}
