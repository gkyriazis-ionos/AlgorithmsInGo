package AlgorithmsWithGo

import "sort"

// BubbleSortInt will sort a list of integers using the bubble sort algorithm.
//
// Big O: O(N^2), where N is the size of the list
func BubbleSortInt(list []int) {
	for i := 0; i < len(list); i++ {
		for j := 0; j < len(list)-1-i; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
			}
		}
	}
}

// BubbleSortString is a bubble sort for string slices. Try implementing it for
// practice.
func BubbleSortString(list []string) {
}

type Person struct {
	age       int
	firstName string
	lastName  string
}

type PersonList []Person

// BubbleSortPerson uses bubble sort to sort Person slices by: Age, then
// LastName, then FirstName. Try implementing it for practice.
func BubbleSortPerson(people []Person) {
}

// BubbleSort uses the standard library's sort.Interface to sort. Try
// implementing it for practice.
func BubbleSort(list sort.Interface) {
	for i := 0; i < list.Len(); i++ {
		for j := 0; j < list.Len()-1-i; j++ {
			if list.Less(j+1, j) {
				list.Swap(j, j+1)
			}
		}
	}
}

func (people PersonList) Less(i, j int) bool {
	if people[i].age < people[j].age {
		return true
	} else if people[i].age > people[j].age {
		return false
	} else {
		if people[i].lastName < people[j].lastName {
			return true
		} else if people[i].lastName > people[j].lastName {
			return false
		} else {
			if people[i].firstName < people[j].firstName {
				return true
			} else if people[i].firstName > people[j].firstName {
				return false
			}
		}
	}
	return false
}

func (people PersonList) Swap(i, j int) {
	people[i], people[j] = people[j], people[i]
}

func (people PersonList) Len() int {
	return len(people)
}
