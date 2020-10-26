package sort

// 遍历一次获取一个数组中的最大值和最小值
func getMaxAndMin(array []int) (max int, min int) {
	max = array[0]
	min = array[0]
	for _, v := range array {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	return
}

func CountSort(array []int) {
	arrayLen := len(array)
	if arrayLen < 1 {
		return
	}
	max, min := getMaxAndMin(array)
	// 初始化统计数组
	countArray := make([]int, max-min+1)

	//遍历原始数组,对统计数组进行赋值
	for _, v := range array {
		countArray[v-min]++ //每出现相同的元素，同样位置进行值+1
	}

	//进行统计数组顺序累加
	for i := 1; i < max-min+1; i++ {
		countArray[i] += countArray[i-1]
	}

	// 初始化排序后数组B
	sortArray := make([]int, arrayLen)
	// 原始数组倒序遍历，查找在统计数组中的值(代表在排序后数组的索引位置)
	for i := arrayLen - 1; i >= 0; i-- {
		v := array[i] - min
		index := countArray[v] - 1 //表示小于等于array[i]元素的个数总和有多少个,也即代表在新数组中的索引(表示索引记得-1)
		sortArray[index] = array[i]
		countArray[v]--
	}

	//复制B数组到原数组A中
	for i := 0; i < arrayLen; i++ {
		array[i] = sortArray[i]
	}
	return
}
