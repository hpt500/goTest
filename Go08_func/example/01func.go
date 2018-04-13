package example

import (
	. "fmt"
)

// 返回一个数
func ReturnMax(x, y int) int {
	Printf("一、1返回%d和%d中最大的数值为\n", x, y)
	var result int
	if x > y {
		result = x
	} else {
		result = y
	}
	return result
}

// 返回多个数->重点是return_type需多个
func ReturnMore(x, y string) (string, string) {
	Printf("一、2返回多个数值\n")
	return x, y
}
