package main

import (
	. "fmt"

	"./example"
)

func main() {
	Println("函数的使用")
	Println()
	// 一、1返回两参数中最大一个的函数
	Printf("返回%d和%d中最大的数值为%d\n", 100, 200, example.ReturnMax(100, 200))
	// 一、2返回多个数值
	x, y := example.ReturnMore("xu", "yi")
	Printf("返回%s和%s等两个字符串\n", x, y)
	Println()

	a, b := 100, 200
	Println("a和b的值分别为:", a, b)
	// 二、1通过值传递交换两个变量的值->不会影响到实际参数
	example.SwapVal(a, b)
	Println("值传递后a和b的值分别为:", a, b)
	// 二、2通过引用传递交换两个变量的实际地址->影响到实际参数
	example.SwapRep(&a, &b)
	Println("引用传递后a和b的值分别为:", a, b)
	Println()

	// 三、1闭包函数的简单使用-->不会影响重新调用函数时的i初始值
	cf := example.Closure1() //返回一个闭包函数 含有i = 0
	Println("第一次调用,返回值为:", cf())
	Println("第二次调用,返回值为:", cf())
	Println("第三次调用,返回值为:", cf())
	// 三、2带参数的闭包函数的简单使用
	cg := example.Closure2(2, 3)
	Println(cg())
	Println(cg())
	Println(cg())
	Println()

	// 四、go语言中同时有函数和方法
	example.FuncMethod()
	Println()
}
