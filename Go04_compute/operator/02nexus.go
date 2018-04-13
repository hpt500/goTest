package operator

import (
	. "fmt"
)

func init() {
	// 二、关系运算符
	var a, b int = 20, 21
	Println("二、关系运算符")
	if a == b {
		Println("第一行a==b输出true")
	} else {
		Println("第一行a==b输出false")
	}
	if a > b {
		Println("第二行a>b输出true")
	} else {
		Println("第二行a>b输出false")
	}
	if a < b {
		Println("第三行a<b输出true")
	} else {
		Println("第三行a<b输出false")
	}
	Println()
}
