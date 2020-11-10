package main

import "fmt"

// 计数排序
// 计数排序(Counting sort) 是一种稳定的线性时间排序算法.计数排序的核心在于将输入的数据值转化为键存储在额外开辟的数组空间中。作为一种线性时间复杂度的排序，计数排序要求输入的数据必须是有确定范围的整数。计数排序使用一个额外的数组C,其中第i个元素是待排序数组A中值等于i的元素的个数。然后根据数组C 来将A中的元素排到正确的位置。

// 算法原理
// 计数排序的算法原理：

// 找出待排序的数组中最大和最小的元素.
// 统计数组中每个值为i的元素出现的次数，存入数组C的第i项.
// 对所有的计数累加(从C中的第一个元素开始，每一项和前一项相加).
// 反向填充目标数组：将每个元素 i放在新数组的第C[i]项，每放一个元素就将C[i]减去1.

func countingSort(arr []int, maxValue int) []int {
	bucketLen := maxValue + 1
	bucket := make([]int, bucketLen) // 初始为0的数组
	sortedIndex := 0
	length := len(arr)
	for i := 0; i < length; i++ {
		bucket[arr[i]] += 1
	}
	for j := 0; j < bucketLen; j++ {
		for bucket[j] > 0 {
			arr[sortedIndex] = j
			sortedIndex += 1
			bucket[j] -= 1
		}
	}
	return arr
}

func main1111() {
	array := []int{69, 16, 48, 2, 2, 3, 32, 10, 27, 17, 42, 29, 8, 28, 12, 9}
	countingSort(array, array[0])
	fmt.Println("BucketSort:", array)
}
