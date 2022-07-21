package AlgorithmsWithGo

import (
	"strconv"
)

// Transposes a number in decimal(dec) to any base required

func DecToBase(dec, base int) string {
	output := ""
	for dec != 0 {
		digit := ""
		switch {
		case dec%base >= 10:
			a := dec % base
			b := a % 10
			c := int('A') + b
			digit = string(rune(c))
		default:
			digit = strconv.Itoa(dec % base)
		}
		output = digit + output
		dec = dec / base
	}
	return output
}
