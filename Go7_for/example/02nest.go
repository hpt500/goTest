package example

import (
	. "fmt"
)

func Nest() {
	Println("二、循环嵌套->输出100以内的素数")
	for i := 2; i < 100; i++ {
		var flag = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
			}
		}
		if flag == true {
			Printf("%d是素数\n", i)
		}
	}
	Println()
}
