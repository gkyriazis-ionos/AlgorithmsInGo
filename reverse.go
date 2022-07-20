package AlgorithmsWithGo

import (
	"strings"
)

func Reverse(word string) string {
	// should I also code the Split word function??
	chars := strings.Split(word, "")
	out_string := ""
	for i := len(chars) - 1; i >= 0; i-- {
		out_string += chars[i]
	}

	return out_string
}
