package util_3gpp

import (
	"fmt"
	"free5gc-cli/lib/openapi/models"
)

func SNssaiToString(snssai *models.Snssai) (str string) {
	if snssai.Sd == "" {
		return fmt.Sprintf("%d-%s", snssai.Sst, snssai.Sd)
	}
	return fmt.Sprintf("%d", snssai.Sst)
}
