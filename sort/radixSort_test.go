package sort

import "testing"

func TestRadixSort(t *testing.T) {
	//var array = []int{329, 457, 657, 839, 436, 720, 355}
	var array = []int{0, 23, 12345, 345, 4569, 234567, 456, 9999, 2}
	RadixSort(array)
	t.Logf("RadixSort result is:%+v\n", array)
	return
}
