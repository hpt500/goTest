package example

import (
	. "fmt"
)

func SwapVal(x, y int) int {
	Println("二、1通过值传递交换两个变量的值->不会影响到实际参数")
	var temp int
	temp = x
	x = y
	y = temp
	return temp
}

func SwapRep(x *int, y *int) int {
	Println("二、2通过引用传递交换两个变量的实际地址->影响到实际参数")
	var temp int
	temp = *x
	*x = *y
	*y = temp
	return temp
}
