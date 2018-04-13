package example

import (
	. "fmt"
)

func ArrFunc(arr []int, size int) float32 {
	Println("一.向函数传递数组并求取数组总和的平均值")
	var i, sum int
	var avg float32
	for i = 0; i < size; i++ {
		sum += arr[i]
	}

	avg = float32(sum) / float32(size)

	return avg

}
