package sort

import (
	"testing"
)

func TestInsertionSort(t *testing.T) {
	lst := []int{2, 4, 1, 5, 6}
	sorted := InsertionSort(lst)
	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] > sorted[i+1] {
			t.Error("list is not sorted!")
		}
	}
}

func TestMerge(t *testing.T) {
	lst := []int{2, 1, 4}
	sorted := MergeSort(lst, 1, len(lst))
	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] > sorted[i+1] {
			t.Error("list is not sorted!")
		}
	}
}
