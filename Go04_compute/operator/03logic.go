package operator

import (
	. "fmt"
)

func init() {
	// 三、逻辑运算符
	var a, b bool = true, false
	Println("三、逻辑运算符")
	if a && b {
		Println("第一行a&&b输出true")
	} else {
		Println("第一行a&&b输出false")
	}
	if a || b {
		Println("第二行a||b输出true")
	} else {
		Println("第二行a||b输出false")
	}
	Println()
}
