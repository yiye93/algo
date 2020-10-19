package sort

// 冒泡升序排序
func BubbuleSortAsc(array []int) {
	arrayLen := len(array)
	lastSwap := arrayLen - 1
	lastSwapTemp := arrayLen - 1

	endFlag := true
	for i := 0; i < arrayLen-1; i++ {
		lastSwap = lastSwapTemp
		for j := 0; j < lastSwap; j++ {
			if array[j] > array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				lastSwapTemp = j
				endFlag = false
			}
		}
		if lastSwap == lastSwapTemp {
			break
		}
		if endFlag {
			break
		}
	}
}

// 冒泡降序排序
func BubbuleSortDesc(array []int) {
	arrayLen := len(array)
	lastSwap := arrayLen - 1
	lastSwapTemp := arrayLen - 1

	endFlag := true
	for i := 0; i < arrayLen-1; i++ {
		lastSwap = lastSwapTemp
		for j := 0; j < lastSwap; j++ {
			if array[j] < array[j+1] {
				array[j], array[j+1] = array[j+1], array[j]
				lastSwapTemp = j
				endFlag = false
			}
		}
		if lastSwap == lastSwapTemp {
			break
		}
		if endFlag {
			break
		}
	}
}
