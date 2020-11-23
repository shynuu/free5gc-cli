package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type SerialNumber struct {
	Value aper.BitString `aper:"sizeLB:16,sizeUB:16"`
}
