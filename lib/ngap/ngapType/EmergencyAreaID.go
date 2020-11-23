package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type EmergencyAreaID struct {
	Value aper.OctetString `aper:"sizeLB:3,sizeUB:3"`
}
