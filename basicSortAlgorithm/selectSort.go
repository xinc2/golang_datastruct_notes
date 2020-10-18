/*
选择排序
	选择排序(Selection sort) 是一种简单直观的排序算法。
	它的工作原理如下:
		首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置，
		然后，再从剩余未排序元素中继续寻找最小（大）元素，
		然后放到已排序序列的末尾。以此类推，直到所有元素均排序完毕。
	选择排序的主要优点与数据移动有关。
		如果某个元素位于正确的最终位置上，则它不会被移动。
		选择排序每次交换一对元素，它们当中至少有一个将被移到其最终位置上，因此对n个元素的表进行排序总共进行至多n-1次交换。
		在所有的完全依靠交换去移动元素的排序方法中，选择排序属于非常好的一种。

算法原理
	选择排序算法的原理如下：
		1. 首先在未排序序列中找到最小（大）元素，存放到排序序列的起始位置
		2. 再从剩余未排序元素中继续寻找最小（大）元素，然后放到已排序序列的末尾。
		重复第二步，直到所有元素均排序完毕。
*/

package main

import (
	"fmt"
)

// SelectionSort 选择排序算法实现
func SelectionSort(arr []int) []int {
	length := len(arr)
	if length < 2 {
		return arr
	}
	for i := 0; i < length-1; i++ {
		minIndex := i
		for j := i + 1; j < length; j++ {
			if arr[j] < arr[minIndex] {
				minIndex = j
			}
		}
		if i != minIndex {
			arr[minIndex], arr[i] = arr[i], arr[minIndex]
		}
		fmt.Println(arr)
	}
	return arr
}

func main11() {
	var arr = []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(SelectionSort(arr))
}
