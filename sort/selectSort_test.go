package sort

import "testing"

func TestSelectSort(t *testing.T) {
	var array = []int{52, 23, 18, 84, 58, 11}
	SelectSort(array)
	t.Logf("SelectSort result is:%+v\n", array)

	array = []int{52, 23, 18, 84, 58, 11, 24}
	SelectSortOpt1(array)
	t.Logf("SelectSortOpt1 result is:%+v\n", array)
}
