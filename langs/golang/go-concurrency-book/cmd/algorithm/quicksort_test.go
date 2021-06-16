package algorithm

import "testing"

func TestQuickSort(t *testing.T) {
	arr := []int{1, 4, 6, 9, 10, 2, 3, 5, 8, 7}
	QuickSort(arr)
}
