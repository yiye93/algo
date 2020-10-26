package sort

// 获取数组的最大位数(即最大值的位数)
func getMaxBit(array []int) (bit int) {
	max := array[0]
	for _, v := range array {
		if v > max {
			max = v
		}
	}
	if max == 0 {
		return 1
	}
	for max != 0 {
		bit++
		max = max / 10
	}
	return bit
}

// 获取指定元素和指定进制位的值(x=12345,b=1,则表示获取元素120,十进制位对应的值,为2)
func getBitValue(x int, bit int) (bitValue int) {
	for i := 1; i <= bit; i++ {
		x = x / 10
	}
	bitValue = x % 10
	return bitValue
}

func RadixSort(array []int) {
	arrayLen := len(array)
	if arrayLen < 1 {
		return
	}
	//获取最大位数
	maxBit := getMaxBit(array)

	// 循环比较每一个进制位
	for i := 0; i < maxBit; i++ {
		//定义10个桶子(每个进制位的值的范围只会在0-9之间)
		var bucketArray = [10][]int{}
		for _, v := range array {
			// 获取到数组元素此时进制位的值
			bitValue := getBitValue(v, i)
			// 将该数组元素加入到对应的bucket中
			bucketArray[bitValue] = append(bucketArray[bitValue], v)
		}
		//array数组重置为依据bucket得到的最新数组
		tmpArray := make([]int, 0)
		// 遍历0-9下标的桶,每个桶存储了该进制的数值对应的数组值
		for _, bucket := range bucketArray {
			for _, v := range bucket {
				tmpArray = append(tmpArray, v)
			}
		}
		copy(array, tmpArray)
	}
	return
}
