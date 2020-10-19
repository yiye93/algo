package sort

import "testing"

func TestShellSort(t *testing.T) {
	var array = []int{84, 83, 88, 87, 61, 50, 70, 60, 80, 99}
	ShellSort(array)
	t.Logf("ShellSort result is:%+v\n", array)
}
