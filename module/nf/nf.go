package nf

import (
	"free5gc-cli/module/nf/api"
)

var nf *NF

type NF struct {
}

func (nf *NF) DropDatabase(db string) error {
	api.Drop(db)
	return nil
}

func (nf *NF) FlushDatabase() error {
	api.Flush()
	return nil
}
