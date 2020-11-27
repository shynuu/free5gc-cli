package api

import (
	"free5gc-cli/lib/openapi/models"
	"strconv"
)

func convertSnssai(snssai string) *models.Snssai {

	ssts := snssai[:2]
	sd := snssai[2:]

	sst64, _ := strconv.Atoi(ssts)
	sst := int32(sst64)

	return &models.Snssai{Sst: sst, Sd: sd}
}
