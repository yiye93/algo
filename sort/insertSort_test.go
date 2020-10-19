package sort

import "testing"

func TestInsertSort(t *testing.T) {
	var array = []int{52, 23, 18, 84, 58, 11}
	//var array = []int{6, 5, 4, 3, 2, 1}
	InsertSort(array)
	t.Logf("InsertSort result is:%+v\n", array)
}
