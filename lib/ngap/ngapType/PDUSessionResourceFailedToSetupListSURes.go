package ngapType

// Need to import "free5gc-cli/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct PDUSessionResourceFailedToSetupListSURes */
/* PDUSessionResourceFailedToSetupItemSURes */
type PDUSessionResourceFailedToSetupListSURes struct {
	List []PDUSessionResourceFailedToSetupItemSURes `aper:"valueExt,sizeLB:1,sizeUB:256"`
}