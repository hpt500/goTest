package main

import (
	. "fmt"

	"./example"
)

func main() {
	Println()
	// 求取数组总和的平均值
	var arr = []int{1000, 2, 3, 17, 50}
	var avg float32 = example.ArrFunc(arr, 5)
	Printf("arr数组总和的平均值为:%f\n", avg)
}
