package main

import (
	. "fmt"

	"./example"
)

var ( // 这种因式分解关键字的写法一般用于声明全局变量
	x int
	y bool
)
var a = "许依衔"
var b string = "565420423@qq.com"
var c bool

// 可进行多变量同类型定义
var f, g, h int

// 可进行多变量在同一行进行赋值
var j, k, v = 1, 7, "ho"

func main() {
	// 这种不带声明格式的只能在函数体中出现
	d := "hello"
	// 在函数体内单纯地给 e 赋值也是不够的，这个值必须被使用
	e := "world"
	Println(x, y, a, b, c, d, e)
	Println()

	// 指针应用
	example.VarPoint()

	// 指针数组
	example.PtrArr()

	// 指向指针的指针变量
	example.PtrToPtr()

	// 另额外说明 若向函数传递指针 则需在函数定义的参数设置为指针类型即可
	// 且其函数内部的指针变量操作会影响实际参数

}
