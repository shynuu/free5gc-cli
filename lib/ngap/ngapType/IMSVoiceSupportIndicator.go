package ngapType

import "free5gc-cli/lib/aper"

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

const (
	IMSVoiceSupportIndicatorPresentSupported    aper.Enumerated = 0
	IMSVoiceSupportIndicatorPresentNotSupported aper.Enumerated = 1
)

type IMSVoiceSupportIndicator struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
