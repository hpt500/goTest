package example

import (
	. "fmt"
)

func Test() {
	Println("三.额外测试--1.简单切片测试")
	sliceTest()
	Println()
	Println("三.额外测试--1.二位数组测试函数")
	twoDimensionArray()
	Println()
}

func sliceTest() {
	arr := []int{1, 2, 3, 4, 5}
	s := arr[:]
	for e := range s {
		Println(s[e])
	}
	s1 := make([]int, 3)
	for e := range s1 {
		Println(s1[e])
	}
}

func twoDimensionArray() {
	/* 数组 - 5 行 2 列*/
	var a = [][]int{{0, 0}, {1, 2}, {2}, {3, 6}, {4, 8}}
	var i, j int

	/* 输出数组元素 */
	for i = 0; i < len(a); i++ {
		for j = 0; j < len(a[i]); j++ {
			Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
