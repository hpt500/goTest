package main

import (
	. "fmt"

	"./example"
)

func main() {
	Println("go语言结构体的应用")
	// 在结构体中我们可以为不同项定义不同的数据类型
	// 结构体定义需要使用 type 和 struct 语句
	// 一旦定义了结构体类型，它就能用于变量的声明，语法格式如下：
	// variable_name := structure_variable_type {value1, value2...valuen}

	// 一.结构体的定义
	example.Struct()

	// 二.结构体指针的应用
	example.StructPtr()

}
