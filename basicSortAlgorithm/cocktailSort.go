package main

import (
	"fmt"
	"math/rand"
	"time"
)

// CocktailSort 鸡尾酒排序，又称双冒泡排序
func CocktailSort(list []int) []int {

	length := len(list)
	for loop := 0; loop <= length/2; loop++ {
		var sorted = false
		var j int = 0
		// find the max one and move it to the end of the line
		for j = 0; j < length-loop-1; j++ {
			if list[j] > list[j+1] {
				list[j], list[j+1] = list[j+1], list[j]
				sorted = true
			}
		}

		if !sorted {
			break
		}

		//find small one
		for ; j > loop; j-- {
			if list[j] < list[j-1] {
				list[j], list[j-1] = list[j-1], list[j]
				sorted = true
			}
		}

		if !sorted {
			break
		}
		fmt.Println(loop, "---->", list)
	}
	return list
}

func main() {
	var length = 10
	var list []int

	// 以时间戳为种子生成随机数，保证每次运行数据不重复
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		list = append(list, int(r.Intn(100)))
	}
	fmt.Println(list)
	CocktailSort(list)
}
