package main

import (
	"fmt"
)

// HeapSortMax used to find the max value in the front of the array
func heapSortMax(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1 // 二叉树的深度；也就是有多少个父节点
	for i := depth; i >= 0; i-- {
		topmax := i
		leftchild := 2*i + 1
		rightchild := 2*i + 2
		if leftchild <= length-1 && arr[leftchild] > arr[topmax] {
			topmax = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] > arr[topmax] {
			topmax = rightchild
		}
		if i != topmax {
			arr[i], arr[topmax] = arr[topmax], arr[i]
		}
	}

	return arr
}

// HeapSortMin used to find the min value in the front of the array
func heapSortMin(arr []int, length int) []int {
	if length <= 1 {
		return arr
	}
	depth := length/2 - 1 // 二叉树的深度；也就是有多少个父节点
	for i := depth; i >= 0; i-- {
		topmin := i
		leftchild := 2*i + 1
		rightchild := 2*i + 2
		if leftchild <= length-1 && arr[leftchild] < arr[topmin] {
			topmin = leftchild
		}
		if rightchild <= length-1 && arr[rightchild] < arr[topmin] {
			topmin = rightchild
		}
		if i != topmin {
			arr[i], arr[topmin] = arr[topmin], arr[i]
		}
	}

	return arr
}

// HeapSortIncr used to sort the heap in ascend order
// swithc the first and last number in each loop
func HeapSortIncr(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastLen := length - i
		heapSortMax(arr, lastLen)
		arr[0], arr[lastLen-1] = arr[lastLen-1], arr[0]
	}

}

// HeapSortDec used to sort the heap in decrement order
// switch the first and last number in each loop
func HeapSortDec(arr []int) {
	length := len(arr)
	for i := 0; i < length; i++ {
		lastLen := length - i
		heapSortMin(arr, lastLen)
		arr[0], arr[lastLen-1] = arr[lastLen-1], arr[0]
	}

}

func main1111() {
	var arr1 = []int{1, 3, 7, 6, 9, 8, 2, 4, 5}
	fmt.Println(arr1)
	HeapSortIncr(arr1)
	fmt.Println(arr1)
	HeapSortDec(arr1)
	fmt.Println(arr1)

}
