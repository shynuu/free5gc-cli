package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type EUTRAencryptionAlgorithms struct {
	Value aper.BitString `aper:"sizeExt,sizeLB:16,sizeUB:16"`
}
