package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type TransportLayerAddress struct {
	Value aper.BitString `aper:"sizeExt,sizeLB:1,sizeUB:160"`
}
