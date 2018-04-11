package example

import (
	. "fmt"
)

func Closure1() func() int {
	Println("三、1闭包函数的简单使用->不会影响重新调用函数时的i初始值")
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func Closure2(x, y int) func() (int, int) {
	Println("三、2带参数的闭包函数的简单使用")
	i := 0
	return func() (int, int) {
		i += 1
		return i, x + y
	}
}
