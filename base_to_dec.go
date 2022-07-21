package AlgorithmsWithGo

import (
	"math"
	"strconv"
	"strings"
)

func BaseToDec(value string, base int) int {
	power := 0
	num := 0
	digits := strings.Split(value, "")
	for i := len(digits) - 1; i >= 0; i-- {
		hexToDec := map[string]string{
			"A": "10",
			"B": "11",
			"C": "12",
			"D": "13",
			"E": "14",
			"F": "15",
		}
		val, ok := hexToDec[digits[i]]
		d := 0
		if ok {
			d, _ = strconv.Atoi(val)
		} else {
			d, _ = strconv.Atoi(digits[i])
		}
		num += d * int(math.Pow(float64(base), float64(power)))
		power += 1
	}
	return num
}
