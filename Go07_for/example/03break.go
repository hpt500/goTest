package example

import (
	. "fmt"
)

func Break() {
	Println("三、break语句的使用->跳出当前循环")
	var a int = 10
	for a < 15 {
		a++
		Printf("当前a的值为%d\n", a)
		if a > 13 {
			break
		}
	}
	Println()
}
