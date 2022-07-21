package AlgorithmsWithGo

import "sort"

// InsertionSortInt will sort a list of integers using the insertion sort
// algorithm.
//
// Big O (without binary search): O(N^2), where N is the size of the list.
// Big O (with binary search): O(N log N), where N is the size of the list.
//
// Test with: go test -run InsertionSortInt$
// The '$' at the end will ensure that the InsertionSortInterface tests won't be run.
func InsertionSortInt(list []int) {
	var sorted []int
	for _, item := range list {
		sorted = insert(sorted, item)
	}

	for i, val := range sorted {
		list[i] = val
	}

}

func insert(sorted []int, item int) []int {
	for i, sortedItem := range sorted {
		if item < sortedItem {
			endList := append([]int{item}, sorted[i:]...)
			sorted = append(sorted[:i], endList...)
			return sorted
		}
	}
	return append(sorted, item)
}

// InsertionSortString uses insertion sort to sort string slices. Try
// implementing it for practice.
func InsertionSortString(list []string) {
}

// InsertionSortPerson uses insertion sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func InsertionSortPerson(people []Person) {
}

// InsertionSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice, but be wary that this particular algorithm is a
// little tricky to implement this way.
func InsertionSort(list sort.Interface) {
	for i := 0; i < list.Len(); i++ {
		for j := 0; j < i; j++ {
			if list.Less(i, j) {
				// index i serves as a "buffer" - a place we can store an "unsorted"
				// item. In this next loop we use it to insert our new item and then
				// rotate every item right one spot in the final list.
				for k := j; k < i; k++ {
					list.Swap(k, i)
				}
			}
		}
	}
}
