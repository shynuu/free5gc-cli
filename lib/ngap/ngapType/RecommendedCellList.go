package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct RecommendedCellList */
/* RecommendedCellItem */
type RecommendedCellList struct {
	List []RecommendedCellItem `aper:"valueExt,sizeLB:1,sizeUB:16"`
}
