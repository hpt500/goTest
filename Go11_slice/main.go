package main

import (
	. "fmt"

	"./example"
)

func main() {
	Println("go语言切片slice的应用")
	Println()
	// 一.定义slice切片并使用len()和cap()
	example.Slice()

	// 二.append()和copy()函数的应用
	example.AppendCopy()

	// 三.额外测试

	example.Test()

}
