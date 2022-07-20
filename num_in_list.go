package AlgorithmsWithGo

func NumInList(list []int, num int) bool {

	for i := 0; i < len(list); i++ {
		if num == list[i] {
			return true
		}
	}
	return false
}
