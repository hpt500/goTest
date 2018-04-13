package example

import (
	. "fmt"
)

func PtrToPtr() {
	Println("三.指向指针的指针变量的应用")
	var a int = 10
	var ptr *int
	var pptr **int // 指向指针的指针变量的声明格式

	ptr = &a //指向变量a的地址

	pptr = &ptr //指向指针变量ptr的地址

	// 获取pptr的值

	Printf("变量a的值为%d\n", a)
	Printf("指正变量ptr的值为%d\n", *ptr)
	Printf("指向指针的指针变量pptr的值为%d\n", **pptr)

	Println()
}
