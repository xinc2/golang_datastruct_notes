/*
插入排序
	插入排序(Insertion Sort) 是一种简单直观的排序算法。它的工作原理是通过构建有序序列，
	对于未排序数据，在已排序序列中从后向前扫描，找到相应位置并插入。
	插入排序在实现上，通常采用in-place排序（即只需用到O(1)的额外空间的排序），
	因而在从后向前扫描过程中，需要反复把已排序元素逐步向后挪位，为最新元素提供插入空间。

算法原理
	插入排序算法原理：
	- 将第一待排序序列第一个元素看做一个有序序列，把第二个元素到最后一个元素当成是未排序序列。
	- 从头到尾依次扫描未排序序列，将扫描到的每个元素插入有序序列的适当位置。
		（如果待插入的元素与有序序列中的某个元素相等，则将待插入元素插入到相等元素的后面。）

如果比较操作的代价比交换操作大的话，可以采用二分查找法来减少比较操作的数目。
该算法可以认为是插入排序的一个变种，称为二分查找插入排序。
*/

package main

import (
	"fmt"
)

// InsertSort 插入排序算法
func InsertSort(arr []int) []int {
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

func main10() {
	var arr = []int{3, 1, 2, 7, 5, 9, 4, 2, 3, 8}
	fmt.Println(InsertSort(arr))
}
