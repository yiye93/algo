package sort

import "testing"

func TestBubbleSort(t *testing.T) {
	var array = []int{52, 23, 18, 84, 58, 11}
	BubbuleSortAsc(array)
	t.Logf("BubbuleSortAsc result is:%+v\n", array)

	array = []int{52, 23, 18, 84, 58, 11}
	BubbuleSortDesc(array)
	t.Logf("BubbuleSortDesc result is:%+v\n", array)
}
