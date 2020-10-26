package sort

// 基础快速排序
func QuickSort(array []int) {
	arrLen := len(array)
	if arrLen < 1 {
		panic("array len is less than 1")
	}
	seperatorSort(array, 0, arrLen-1)
	return
}

// 小数组序列使用插入排序优化版
func QuickSortOpt1(array []int) {
	arrLen := len(array)
	if arrLen < 1 {
		panic("array len is less than 1")
	}
	seperatorSortOpt1(array, 0, arrLen-1)
	return
}

// 三向切分法优化版
func QuickSortOpt2(array []int) {
	arrLen := len(array)
	if arrLen < 1 {
		panic("array len is less than 1")
	}
	seperatorSortOpt2(array, 0, arrLen-1)
	return
}

func seperatorSort(array []int, start int, end int) {
	if start >= end {
		return
	}
	// 找到分区索引,将该区间分为两段
	pos := patition(array, start, end)
	// 左分区排序
	seperatorSort(array, start, pos-1)
	// 右分区排序
	seperatorSort(array, pos+1, end)
	return
}

func seperatorSortOpt1(array []int, start int, end int) {
	if start >= end {
		return
	}
	// 这里当序列<=5时使用插入排序
	if end-start <= 5 {
		InsertSort(array[start : end+1])
		return
	}
	// 找到分区索引,将该区间分为两段
	pos := patition(array, start, end)
	// 左分区排序
	seperatorSortOpt1(array, start, pos-1)
	// 右分区排序
	seperatorSortOpt1(array, pos+1, end)
	return
}

//在lt之前的(start~lt-1)都小于中间值
//在gt之前的(gt+1~end)都大于中间值
//在lt~i-1的都等于中间值
func seperatorSortOpt2(array []int, start int, end int) {
	if start >= end {
		return
	}
	priv := array[start]
	lt := start
	i := start + 1
	gt := end
	for i <= gt {
		if array[i] < priv {
			array[i], array[lt] = array[lt], array[i]
			i++
			lt++
		} else if array[i] > priv {
			array[i], array[gt] = array[gt], array[i]
			gt--
		} else {
			i++
		}
	}
	seperatorSortOpt2(array, start, lt-1)
	seperatorSortOpt2(array, gt+1, end)
	return
}

func patition(array []int, start int, end int) (pos int) {
	priv := array[start]
	i := start //哨兵i
	j := end   //哨兵j
	for i < j {
		// 从右往左扫描,找到小于基准元素的值
		for array[j] >= priv && i < j {
			j--
		}
		// 从左往右扫描,找到大于基准元素的值
		for array[i] <= priv && i < j {
			i++
		}
		if i < j {
			//哨兵相遇
			array[i], array[j] = array[j], array[i]
		}
	}
	// 交换基准元素和哨兵,使得交换后基准元素左边的都比它小,右边都比它大
	array[i], array[start] = array[start], array[i]
	return i
}
