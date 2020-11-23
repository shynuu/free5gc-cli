package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type NGRANTraceID struct {
	Value aper.OctetString `aper:"sizeLB:8,sizeUB:8"`
}
