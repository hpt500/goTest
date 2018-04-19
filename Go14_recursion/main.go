package main

import (
	. "fmt"

	"./example"
)

func main() {
	Println("go语言递归函数的应用")
	Println()
	// 一.递归函数中的阶乘应用
	example.Factorial()
	// 二.斐波那契数列
	example.Fibonacci()

}
