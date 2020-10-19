package sort

import "testing"

func TestMergeSort(t *testing.T) {
	var array = []int{8, 4, 5, 7, 1, 3, 6, 2}
	MergeSort(array)
	t.Logf("MergeSort result is:%v", array)
}
