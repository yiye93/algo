package sort

func MergeSort(array []int) {
	arrayLen := len(array)
	if arrayLen < 1 {
		panic("array len is less than 1")
	}
	start := 0
	end := arrayLen - 1
	mergeSort(array, start, end)
}

func mergeSort(array []int, start int, end int) {
	if start >= end {
		return
	}
	//找到数组中的中间点，把数组分为两部分
	mid := (start + end) / 2
	//对数组的左部分调用mergeSort 函数
	mergeSort(array, start, mid)
	//对数组的右部分调用mergeSort 函数
	mergeSort(array, mid+1, end)
	//合并数组的左右部分
	merge(array, start, mid, end)
	return
}

func merge(array []int, start int, mid int, end int) {
	// 申请一个辅助数组来合并两个有序数组，这两个数组是 array[start,mid],array[mid+1,end)
	tempArr := make([]int, 0, end-start+1)
	// 左数组移标
	i := start
	// 右数组移标
	j := mid + 1

	k := 0
	// 两个数组开始比较
	for ; i <= mid && j <= end; k++ {
		if array[i] <= array[j] {
			tempArr = append(tempArr, array[i])
			i++
		} else {
			tempArr = append(tempArr, array[j])
			j++
		}
	}

	//剩余序列拼接
	for ; i <= mid; i++ {
		tempArr = append(tempArr, array[i])
	}

	for ; j <= end; j++ {
		tempArr = append(tempArr, array[j])
	}

	// 将辅助数组的元素复制回原数组
	for i, v := range tempArr {
		array[start+i] = v
	}
	return
}
