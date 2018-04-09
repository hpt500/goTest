package condition

import (
	. "fmt"
)

func init() {
	var s int // 声明变量s是需要判断的数
	Println("一、判断输入的数值是否是偶数")
	Println("请输入一个数字:")
	// Scan(&s) // 输入语句
	if s%2 == 0 {
		Println("s 是偶数,其值为: ", s)
	} else {
		Println("s 不是偶数,其值为: ", s)
	}
	Println()
}
