package main

import "fmt"

// 一、常量-->在程序运行时，不会被修改的量
// 常量-->常量中的数据类型只可以是布尔型、数字型（整数型、浮点型和复数）和字符串型
const LENGTH int = 10
const WIDTH int = 5

var alen = 20

var area, area2 int

const a, b, c = 1, false, "str"

// 二、可用作枚举-->数字 0、1 和 2 分别代表未知性别、女性和男性。
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

// 三、iota，特殊常量，可以认为是一个可以被编译器修改的常量
// 在每一个const关键字出现时，被重置为0，然后再下一个const出现之前，每出现一次iota，其所代表的数字会自动增加1。
const (
	i1 = iota //0
	i2 = iota //1
	i3 = iota //2
)

// 可简写为
const (
	i11 = iota //0
	i21        //1
	i31        //2
)

// iota用法指南--注从iota初始定义(0)开始 每出现一行iota值加1
const (
	o1 = iota //0
	o2        //1
	o3        //2
	o4 = "ha" //独立值,但iota仍会加1
	o5        //默认同步上一个定义,iota++
	o6 = 100  //独立值,iota++
	o7        //默认同步上一个定义,iota++
	o8 = iota //7,恢复计数但并不会重置-->同一个const内开始定义-->即可通过重定义iota来恢复计数
	o9        //8
)

// iota实践用法--"<<"符号的运用
const (
	s1 = 1 << iota //十进制数值(1)转成二进制数据并向左移动0位即1
	s2 = 3 << iota //十进制数值(3)转成二进制数据并向左移动1位即110
	s3             //十进制数值(3)转成二进制数据并向左移动2位即1100
	s4             //十进制数值(3)转成二进制数据并向左移动3位即11000
)

func main() {
	alen = 12
	area = LENGTH * WIDTH
	area2 = alen * WIDTH
	fmt.Printf("面积为 : %d", area)
	println()
	fmt.Printf("面积2为 : %d", area)
	println()
	println(a, b, c)
	println(Unknown)

	println()
	println("iota用法")
	println("t1:", i1, ";t2:", i2, ";t3:", i3)
	println("iota简写用法")
	println("t11:", i11, ";t21:", i21, ";t31:", i31)
	println("iota详细用法")
	println(o1, o2, o3, o4, o5, o6, o7, o8, o9)
	println("iota实践用法")
	println(s1, s2, s3, s4)

}
