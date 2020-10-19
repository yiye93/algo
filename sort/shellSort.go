package sort

// 这里采用增量折半策略的希尔排序
func ShellSort(array []int) {
	arrayLen := len(array)
	// 增量折半
	for step := arrayLen / 2; step >= 1; step /= 2 {
		// 分为step组
		for i := 0; i < step; i++ {
			// 演变为步长为step的插入排序
			for j := i + step; j < arrayLen; j += step {
				deal := array[j]
				if deal > array[j-step] {
					continue
				}
				destIndex := j
				for k := j - step; k >= 0 && array[k] > deal; k -= step {
					array[k+step] = array[k]
					destIndex = k
				}
				array[destIndex] = deal
			}
		}
	}
	return
}
