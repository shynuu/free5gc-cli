package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type WarningMessageContents struct {
	Value aper.OctetString `aper:"sizeLB:1,sizeUB:9600"`
}
