package example

import (
	. "fmt"
)

func Goto() {
	Println("五、goto语句的使用->可用来实现条件转移,构成循环,跳出循环体等功能")
	var c int = 10
Loop:
	for c < 15 {
		if c == 12 {
			/* 跳过迭代 */
			c++
			goto Loop
		}
		Printf("c当前的值为%d\n", c)
		c++
	}
	Println()
}
