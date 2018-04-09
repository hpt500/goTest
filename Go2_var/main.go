package main

import "fmt"

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
	fmt.Println(x, y, a, b, c, d, e)
}
