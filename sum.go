package AlgorithmsWithGo

func Sum(numbers []int) int {
	s := 0
	// underscore are the indexes. val are the values.
	for _, val := range numbers {
		s += val
	}
	return s
}
