package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type RATRestrictionInformation struct {
	Value aper.BitString `aper:"sizeExt,sizeLB:8,sizeUB:8"`
}
