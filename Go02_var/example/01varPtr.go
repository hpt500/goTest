package example

import (
	. "fmt"
)

func VarPoint() {
	Println("一.指针变量的应用")
	var a int = 10
	var b *int // 声明指针变量

	b = &a //指针变量的存储地址

	Println("输出a变量的地址:", &a)
	// 指针变量的存储地址
	Println("指针变量b的存储地址:", b)
	// 使用指针访问值
	Println("*b变量的值为:", *b)
	Println()

	Println("一.空指针的应用")
	var ptr *int
	// 当一个指针被声明后但没有分配任何变量时,其值为nil 类似于null等空值
	Println("指针变量ptr的值为:", ptr)
	Printf("ptr的值(16z)为:%x\n", ptr)
	if ptr == nil {
		Println("ptr是空指针")
	} else {
		Println("ptr不是空指针")
	}
	Println()
}
