package sort

func InsertSort(array []int) {
	arrayLen := len(array)
	for i := 1; i < arrayLen; i++ {
		dealElem := array[i]
		// 比较的元素大于已排好序的数组,不需要动
		if dealElem > array[i-1] {
			continue
		}
		destIndex := i
		// 一直往左边找，比待排序大的数都往后挪，腾空位给待排序插入
		for j := i - 1; j >= 0 && dealElem < array[j]; j-- {
			destIndex = j
			array[j+1] = array[j]
		}
		//待排序的数插入空位
		array[destIndex] = dealElem
	}
}
