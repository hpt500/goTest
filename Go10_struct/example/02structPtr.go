package example

import (
	. "fmt"
)

type Test struct {
	score  int
	result string
}

func StructPtr() {
	Println("二.结构体指针变量的应用")
	// 其他指针比较
	var str string
	str = "测试"
	var strPtr *string
	strPtr = &str
	Println("相较于结构体数组对象之类的内存单位->strPtr的值为:", strPtr)

	var test1 Test
	test1.score = 90
	test1.result = "优秀"

	var testPtr *Test
	testPtr = &test1

	Println("Test of score(testPtr.score):", testPtr.score)

	Println()

}
