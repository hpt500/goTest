package example

import (
	. "fmt"
)

func Continue() {
	Println("四、continue语句的使用->跳出当前循环执行下一次循环")
	var b int = 10
	for b < 15 {
		b++
		if b == 12 {
			continue
		}
		Printf("当前b的值为%d\n", b)
	}
	Println()
}
