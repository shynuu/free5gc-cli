package ngapType

import "free5gc/lib/aper"

// Need to import "free5gc/lib/aper" if it uses "aper"

type SD struct {
	Value aper.OctetString `aper:"sizeLB:3,sizeUB:3"`
}
