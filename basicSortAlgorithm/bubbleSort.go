/*
冒泡排序
	冒泡排序(Bubble Sort) 是一种计算机科学领域的较简单的排序算法。
	它重复地走访过要排序的数列，一次比较两个元素，如果他们的顺序错误就把他们交换过来。
	数列遍历的工作是重复地进行直到没有再需要交换，此时该数列就已经排序完成。
	冒泡排序还有一种优化算法，就是立一个 flag，当在一趟序列遍历中元素没有发生交换，则证明该序列已经有序。

算法原理
	冒泡排序算法的原理如下：
		1. 比较相邻的元素。如果第一个比第二个大，就交换他们两个。
		2. 对每一对相邻元素做同样的工作，从开始第一对到结尾的最后一对。在这一点，最后的元素应该会是最大的数。
		3. 针对所有的元素重复以上的步骤，除了最后一个。
		4. 持续每次对越来越少的元素重复上面的步骤，直到没有任何一对数字需要比较。
*/

package main

import "fmt"

// BubbleSort move the biggest/lowest one to the end in each loop
// 适用场景: 将一个值插入到一个有序列表中，进行重新排序时，使用冒泡排序。性能最佳
func BubbleSort(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	return arr
}

func main110() {
	var arr1 = []int{1, 3, 7, 6, 9, 8, 2, 4, 5}
	fmt.Println(arr1)
	BubbleSort(arr1)
	fmt.Println(arr1)
}
