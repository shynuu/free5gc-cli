package main

import "free5gc-cli/freecli"

// load configuration file.yaml
// load ue file.yaml

// func parseSupi(supi string) []uint8 {
// 	var
// 	return nil
// }

// send InitialUeMessage(Registration Request)(imsi- 208 93 00 00 74 87)
func main() {

	// mobileIdentity5GS := nasType.MobileIdentity5GS{
	// 	Len:    12, // suci
	// 	Buffer: []uint8{0x01, 0x02, 0xf8, 0x39, 0xf0, 0xff, 0x00, 0x00, 0x01, 0x00, 0x47, 0x79},
	// }

	freecli.Initialize()
	freecli.Run()

}
