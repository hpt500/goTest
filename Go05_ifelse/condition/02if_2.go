package condition

import (
	. "fmt"
)

func init() {
	// 条件判断语句里面允许声明一个变量，这个变量的作用域只能在该条件逻辑块内，其他地方就不起作用了
	Println("二、if条件判断允许添加变量")
	if num := 9; num < 0 {
		Println("num 的值小于0,其值为: ", num)
	} else if num < 10 {
		Println("num 的值小于10,其值为: ", num)
	} else {
		Println("num 的值大于等于10,其值为: ", num)
	}
	Println()
}
