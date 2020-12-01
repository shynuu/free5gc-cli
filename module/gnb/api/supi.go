package api

import (
	"encoding/hex"
	"fmt"
	"strings"
)

func SHIFT(b byte) byte {
	return b<<4 + b>>4
}

func SHIFTArray(array []byte) []byte {
	buffer := make([]byte, len(array))
	for i, v := range array {
		buffer[i] = SHIFT(v)
	}
	return buffer
}

func parseMCC(mcc string) ([]byte, error) {

	mcc = mcc + "0"
	b0, err := hex.DecodeString(mcc[:2])
	if err != nil {
		return nil, err
	}
	b1, err := hex.DecodeString(mcc[2:])
	if err != nil {
		return nil, err
	}

	return append(SHIFTArray(b0), SHIFTArray(b1)...), nil
}

func parseMNC(mnc string) ([]byte, error) {

	var b0 []byte
	var b1 []byte
	var err error

	if mnc[0] == '0' {
		mnc = "0" + mnc
		b1, err = hex.DecodeString(mnc[:2])
		if err != nil {
			return nil, err
		}
		b0, err = hex.DecodeString(mnc[2:])
		if err != nil {
			return nil, err
		}
		b1[0] = b1[0] + 0b1111
	} else {
		mnc = mnc + "0"
		b0, err = hex.DecodeString(mnc[:2])
		if err != nil {
			return nil, err
		}
		b1, err = hex.DecodeString(mnc[2:])
		if err != nil {
			return nil, err
		}

	}

	return append(SHIFTArray(b0), SHIFTArray(b1)...), nil
}

func formatSupi(supi string) string {
	var s string
	if len(supi)%2 != 0 {
		s = supi[:3] + "0" + supi[3:6] + supi[6:]
	} else {
		s = supi
	}
	return s
}

func parseIdentifier(identifier string) ([]byte, error) {

	b, err := hex.DecodeString(identifier)
	if err != nil {
		return nil, err
	}
	return SHIFTArray(b), nil

}

func supiToSuci(supi string) ([]byte, error) {
	// 	Buffer: []uint8{0x01, 0x02, 0xf8, 0x39, 0xf0, 0xff, 0x00, 0x00, 0x01, 0x00, 0x47, 0x78},
	var buffer = []uint8{0x01, 0x02, 0xf8, 0x39, 0xf0, 0xff, 0x00, 0x00, 0x00, 0x00, 0x47, 0x78}
	fmt.Println(buffer)

	s := strings.Trim(supi, "imsi-")
	s = formatSupi(s)
	mcc := s[:3]
	mnc := s[3:6]

	mccB, err := parseMCC(mcc)
	if err != nil {
		return nil, err
	}
	mncB, err := parseMNC(mnc)
	if err != nil {
		return nil, err
	}

	fmt.Println(mccB, mncB)

	buffer[1] = mccB[0]
	buffer[2] = mccB[1] + mncB[1]
	buffer[3] = mncB[0]

	identifier := s[6:]
	identifierB, err := parseIdentifier(identifier)
	if err != nil {
		return nil, err
	}

	buffer[8] = identifierB[0]
	buffer[9] = identifierB[1]
	buffer[10] = identifierB[2]
	buffer[11] = identifierB[3]
	return buffer, nil
}
