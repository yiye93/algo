package sort

import (
	"fmt"
	"math"
)

// 桶采用数组实现
func BucketSort(array []int) {
	arrayLen := len(array)
	if arrayLen < 1 {
		return
	}

	// 设置桶数量
	bucketNum := 5
	bucketArray := make([][]int, bucketNum)

	//获取数组的最大值和最小值
	max, min := getMaxAndMin(array)
	if max == min {
		return
	}
	// 计算出每个桶的值区间大小,记得+1
	bucketSpace := float64(max-min+1) / float64(bucketNum)
	// 遍历原始数组,按范围把数组元素一次放入对应的桶中
	for _, v := range array {
		// 向下取整
		bucketId := int(math.Floor((float64)(v-min) / float64(bucketSpace)))
		bucketArray[bucketId] = append(bucketArray[bucketId], v)
	}

	//对每个桶采用排序算法进行排序,并且合并
	var sortArray = make([]int, 0)
	for index, bucket := range bucketArray {
		// 这里我们采用插入排序算法
		InsertSort(bucket)
		sortArray = append(sortArray, bucket...)
		fmt.Printf("bucketId:%v | array:%v\n", index, bucket)
	}
	//替换原数组
	copy(array, sortArray)
	return
}

type LinkNode struct {
	Data int
	Next *LinkNode
}

// 桶采用链表实现
func BucketSort1(array []int) {
	arrayLen := len(array)
	if arrayLen < 1 {
		return
	}
	// 设置桶数量,并且初始化每个桶
	bucketNum := 5
	bucketArray := make([]*LinkNode, bucketNum)
	for i := 0; i < bucketNum; i++ {
		bucketArray[i] = &LinkNode{} //初始化头指针
	}

	//获取数组的最大值和最小值
	max, min := getMaxAndMin(array)
	if max == min {
		return
	}
	// 计算出每个桶的值区间大小,记得+1
	bucketSpace := float64(max-min+1) / float64(bucketNum)
	// 遍历原始数组,按范围把数组元素一次放入对应的桶中
	for _, v := range array {
		// 向下取整
		bucketId := int(math.Floor((float64)(v-min) / float64(bucketSpace)))
		linkHead := bucketArray[bucketId] //取出链表头指针
		curNode := linkHead.Next
		curNodePreNode := linkHead
		if linkHead.Next == nil {
			linkHead.Next = &LinkNode{Data: v}
			continue
		}
		//遍历进行比较并插入
		var find = false
		for curNode != nil {
			if v >= curNode.Data {
				curNodePreNode = curNode
				curNode = curNode.Next
				continue
			}
			newNode := &LinkNode{Data: v}
			curNodePreNode.Next = newNode
			newNode.Next = curNode
			find = true
			break
		}
		if !find {
			curNodePreNode.Next = &LinkNode{Data: v}
		}
	}

	//合并桶
	var sortArray = make([]int, 0)
	for _, bucketLink := range bucketArray {
		var tmp []int
		// 遍历链表
		curNode := bucketLink.Next
		for curNode != nil {
			tmp = append(tmp, curNode.Data)
			curNode = curNode.Next
		}
		//fmt.Printf("bucketId:%v | array is:%v\n", bucketId, tmp)
		sortArray = append(sortArray, tmp...)
	}
	copy(array, sortArray)
	return
}
