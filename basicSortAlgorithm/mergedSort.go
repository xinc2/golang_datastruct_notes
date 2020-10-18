/*
归并排序
	归并排序(Merge sort) 是建立在归并操作上的一种有效的排序算法,该算法是采用分治法（Divide and Conquer）的一个非常典型的应用。
	将已有序的子序列合并，得到完全有序的序列；即先使每个子序列有序，再使子序列段间有序。若将两个有序表合并成一个有序表，称为二路归并。

作为一种典型的分而治之思想的算法应用，归并排序的实现由两种方法：
	自上而下的递归(所有递归的方法都可以用迭代重写，所以就有了第 2 种方法).
	自下而上的迭代.
算法原理
	归并排序算法原理：
		申请空间，使其大小为两个已经排序序列之和，该空间用来存放合并后的序列.
		设定两个指针，最初位置分别为两个已经排序序列的起始位置.
		比较两个指针所指向的元素，选择相对小的元素放入到合并空间，并移动指针到下一位置.
		重复上一步直到某一指针达到序列尾；
		将另一序列剩下的所有元素直接复制到合并序列尾.
*/

package main

import (
	"fmt"
)

// MergeSort 合并排序
func MergeSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	n := length / 2
	left := MergeSort(arr[:n])
	right := MergeSort(arr[n:])

	return merge(left, right)
}

// merge 有序的合并两个子序列
func merge(left, right []int) []int {
	tmpSlice := make([]int, 0) // 分配一个临时空间，用来存储合并后的两个的子序列
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			tmpSlice = append(tmpSlice, left[i])
			i++
		} else {
			tmpSlice = append(tmpSlice, right[j])
			j++
		}
	}

	tmpSlice = append(tmpSlice, left[i:]...)
	tmpSlice = append(tmpSlice, right[j:]...)

	return tmpSlice
}

func main111() {
	var slice1 = []int{9, 10, 1, 4, 2, 3, 37, 87, 2, 7, 6, 21}
	slice1 = MergeSort(slice1)
	fmt.Println(slice1)
}
