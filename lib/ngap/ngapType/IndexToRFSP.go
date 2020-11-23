package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

type IndexToRFSP struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:256"`
}
