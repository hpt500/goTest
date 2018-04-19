package example

import (
	. "fmt"
)

func Fibonacci() {
	Println("二.递归函数实现斐波那契数列")
	for i := 0; i < 10; i++ {
		Printf("%d\t", FibFunc(i))
	}
	Println()
}

func FibFunc(n int) int {
	if n < 2 {
		return n
	}
	return FibFunc(n-2) + FibFunc(n-1)

}
