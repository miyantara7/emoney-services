package util

import (
	crand "crypto/rand"
	"math/big"
)

func PadRight(str, pad string, lenght int) string {
	for {
		str += pad
		if len(str) > lenght {
			return str[0:lenght]
		}
	}
}

func GenerateNoKartu() string {
	number0, err := crand.Int(crand.Reader, big.NewInt(9999999999999999))
	if err != nil {
		panic(err)
	}

	return PadRight(number0.String(), "0", 16)
}
