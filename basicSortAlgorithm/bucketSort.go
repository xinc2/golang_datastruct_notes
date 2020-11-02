// 桶排序
// 桶排序(Bucket sort),工作的原理是将数组分到有限数量的桶里。每个桶再个别排序（有可能再使用别的排序算法或是以递归方式继续使用桶排序进行排序）。桶排序是鸽巢排序的一种归纳结果。 桶排序是计数排序的升级版。它利用了函数的映射关系，高效与否的关键就在于这个映射函数的确定。为了使桶排序更加高效，我们需要做到这两点：

// 在额外空间充足的情况下，尽量增大桶的数量.

// 使用的映射函数能够将输入的 N 个数据均匀的分配到 K 个桶中.

// 算法原理
// 桶排序的算法原理：

// 设置一个定量的数组当作空桶子.
// 寻访序列，并且把项目一个一个放到对应的桶子去.
// 对每个不是空的桶子进行排序.
// 从不是空的桶子里把项目再放回原来的序列中.
// 桶排序算法实现:

package main

import "fmt"

// InsertSort 插入排序算法
func InsertSort1(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	for i := 1; i < length; i++ {
		for j := i; j > 0 && arr[j-1] > arr[j]; j-- {
			arr[j-1], arr[j] = arr[j], arr[j-1]
		}
		fmt.Println("InsertSort-> ", arr)
	}
	return arr
}

//获取数组最大值
func getMaxInArr(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

//桶排序
func BucketSort(arr []int) []int {
	// 定义桶数
	num := len(arr)

	// 找到k（数组最大值）
	max := getMaxInArr(arr)

	// 二维切片 (维度1表示桶的编号,维度2是一个数组,存放对应桶的元素)
	buckets := make([][]int, num)

	//分配入桶
	index := 0
	for i := 0; i < num; i++ {
		index = arr[i] * (num - 1) / max // 将待排序的元素分配桶  index = value * (n-1) /k
		buckets[index] = append(buckets[index], arr[i])
	}

	//桶内排序
	tmpPos := 0
	for i := 0; i < num; i++ {
		bucketLen := len(buckets[i])
		if bucketLen > 0 {
			InsertSort1(buckets[i])
			copy(arr[tmpPos:], buckets[i])
			tmpPos += bucketLen
		}
	}

	return arr
}

func main210() {
	array := []int{31, 16, 37, 2, 13, 32, 10, 27, 7, 42, 29, 18, 28, 12, 9}
	BucketSort(array)
	fmt.Println("BucketSort:", array)
}
