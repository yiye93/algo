package sort

import "testing"

func TestCountSort(t *testing.T) {
	var array = []int{16, 19, 31, 103, 19, 25, 19, 25}
	CountSort(array)
	t.Logf("CountSort result is:%+v\n", array)
}
