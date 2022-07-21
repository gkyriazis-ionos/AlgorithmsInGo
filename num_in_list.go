package AlgorithmsWithGo

// NumInList checks if an integer number is contained in a list of integer numbers
func NumInList(list []int, num int) bool {

	for i := 0; i < len(list); i++ {
		if num == list[i] {
			return true
		}
	}
	return false
}
