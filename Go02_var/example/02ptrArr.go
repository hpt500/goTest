package example

import (
	. "fmt"
)

func PtrArr() {
	Println("二.指针数组的应用")
	// 有一种情况，我们可能需要保存数组，这样我们就需要使用到指针
	const MAX int = 3
	a := []int{10, 20, 30}
	// 定义数组时 若需被赋值 则需设置数组的长度*123*
	var b [MAX]*int
	for i := 0; i < MAX; i++ {
		b[i] = &a[i] // 整数地址赋予给指针数组
	}
	for i := 0; i < MAX; i++ {
		Printf("a[%d]=%d\n", i, *b[i])
	}
	Println()

}
