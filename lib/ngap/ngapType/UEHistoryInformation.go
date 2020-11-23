package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct UEHistoryInformation */
/* LastVisitedCellItem */
type UEHistoryInformation struct {
	List []LastVisitedCellItem `aper:"valueExt,sizeLB:1,sizeUB:16"`
}
