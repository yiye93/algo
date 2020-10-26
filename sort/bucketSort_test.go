package sort

import "testing"

func TestBucketSort(t *testing.T) {
	var array = []int{63, 157, 189, 51, 101, 47, 141, 121, 157, 156, 194, 117, 98, 139, 67, 133, 181, 13, 28, 109}
	BucketSort1(array)
	t.Logf("BucketSort result is:%v", array)
	return
}
