package AlgorithmsWithGo

import (
	"AlgorithmsWithGo/sorttest"
	"testing"
)

func TestInsertionSortInt(t *testing.T) {
	sorttest.TestInt(t, InsertionSortInt)
}

func BenchmarkInsertionSortInt(b *testing.B) {
	sorttest.BenchmarkInt(b, InsertionSortInt)
}

func TestInsertionSortString(t *testing.T) {
	sorttest.TestString(t, InsertionSortString)
}

func TestInsertionSortInterface(t *testing.T) {
	sorttest.TestInterface(t, InsertionSort)
}

func TestPersonInsertion(t *testing.T) {
	testPeople(t, InsertionSort)
}
