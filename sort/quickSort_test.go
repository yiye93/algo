package sort

import "testing"

func TestQuickSort(t *testing.T) {
	var array = []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	QuickSort(array)
	t.Logf("QuickSort result is:%v", array)
	array = []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	QuickSortOpt1(array)
	t.Logf("QuickSortOpt1 result is:%v", array)
	array = []int{4, 8, 2, 4, 4, 4, 7, 9}
	QuickSortOpt2(array)
	t.Logf("QuickSortOpt2 result is:%v", array)
}
