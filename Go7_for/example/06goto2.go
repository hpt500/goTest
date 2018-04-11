package example

import (
	. "fmt"
)

func GotoExa() {
	Println("六、goto语句的使用例子->输出100以内的素数")
	var i int = 1
Loop:
	for i < 100 {
		i++
		for j := 2; j < i; j++ {
			if i%j == 0 {
				goto Loop
			}
		}
		Printf("%d是素数\n", i)
	}
	Println()

}
