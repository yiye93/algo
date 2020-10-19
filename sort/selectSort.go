package sort

func SelectSort(array []int) {
	arrayLen := len(array)
	// 进行N-1轮迭代
	for i := 0; i < arrayLen-1; i++ {
		// 最小元素的索引位
		minIndex := i
		// 记录下最小数
		min := array[i]
		for j := i + 1; j < arrayLen; j++ {
			if array[j] < min {
				// 变更最小数和最小数下标
				min = array[j]
				minIndex = j
			}
		}
		// 进行元素位置交换
		if minIndex != i {
			array[i], array[minIndex] = array[minIndex], array[i]
		}
	}
}

func SelectSortOpt1(array []int) {
	arrayLen := len(array)
	// 进行N/2轮迭代
	for i := 0; i < arrayLen/2; i++ {
		// 最小元素的索引位
		minIndex := i
		// 最大元素的索引位
		maxIndex := i
		// 记录下最小数
		min := array[i]
		// 记录下最大数
		max := array[i]

		for j := i + 1; j < arrayLen-i; j++ {
			if array[j] < min {
				// 变更最小数和最小数下标
				min = array[j]
				minIndex = j
			}
			if array[j] > max {
				// 变更最大数和最大数下标
				max = array[j]
				maxIndex = j
			}
		}
		// 进行元素位置交换
		if minIndex != i {
			array[i], array[minIndex] = array[minIndex], array[i]
		}
		if maxIndex != arrayLen-i-1 {
			array[arrayLen-i-1], array[maxIndex] = array[maxIndex], array[arrayLen-i-1]
		}
	}
}
