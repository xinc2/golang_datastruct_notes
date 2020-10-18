/*
快速排序
	快速排序使用分治法（Divide and conquer）策略来把一个串行（list）分为两个子串行（sub-lists）。
	快速排序又是一种分而治之思想在排序算法上的典型应用。本质上来看，快速排序应该算是在冒泡排序基础上的递归分治法。
	快速排序的名字起的是简单，当听到这个名字的时候其实你就知道它存在的意义，就是快速排序，而且效率高！它是处理大数据最快的排序算法之一了。

	快速排序的最坏运行情况是 O(n²)，比如说顺序数列的快排。
	但它的平摊期望时间是 O(nlogn)，且 O(nlogn) 记号中隐含的常数因子很小，比复杂度稳定等于 O(nlogn) 的归并排序要小很多。
	所以，对绝大多数顺序性较弱的随机数列而言，快速排序总是优于归并排序。

算法原理
	快速排序的算法原理:
		从数列中挑出第一个元素，称为 “基准”(pivot).
		重新排序数列，所有元素比基准值小的摆放在基准前面，所有元素比基准值大的摆在基准的后面（相同的数可以到任一边）。
			在这个分区退出之后，该基准就处于数列的中间位置。这个称为分区（partition）操作.
		递归地（recursive）把小于基准值元素的子数列和大于基准值元素的子数列排序.
		递归的最底部情形，是数列的大小是零或一，也就是永远都已经被排序好了。虽然一直递归下去，但是这个算法总会退出，
			因为在每次的迭代（iteration）中，它至少会把一个元素摆到它最后的位置去。
*/

package main

import (
	"fmt"
)

// QuickSortAsending 快速排序
func QuickSortAsending(arr []int, start, end int) {

	if start > end {
		return
	}

	pivot := arr[start] // 基准值
	i, j := start, end

	for i != j {
		// 从最右侧向左开始查找,直到找到小于基准值的数
		for arr[j] >= pivot && i < j {
			j--
		}

		// 从左向右开始查找，直到找到大于基准值的数
		for arr[i] <= pivot && i < j {
			i++
		}

		//交换两个数在数组中的位置
		if i < j {
			arr[i], arr[j] = arr[j], arr[i]
		}

		// i,j相遇,将基准值归位
		arr[start] = arr[i]
		arr[i] = pivot
	}

	// 递归的继续处理基准值左边的数组
	QuickSortAsending(arr, start, i-1)

	// 递归的继续处理基准值右边的数组
	QuickSortAsending(arr, i+1, end)

}

func main() {
	var arr = []int{9, 10, 1, 4, 2, 3, 37, 87, 2, 7, 6, 21}
	QuickSortAsending(arr, 0, len(arr)-1)
	fmt.Println(arr)
}
