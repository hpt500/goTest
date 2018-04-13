package condition

import (
	. "fmt"
)

func init() {
	Println("四、switch判断变量的类型")
	// 除了普遍类似js的switch功能外
	// 还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型
	// 即type switch
	var x interface{}

	switch i := x.(type) {
	case nil:
		Printf("x 的类型为:%T", i)
	case int:
		Printf("x 的类型为:int")
	case float64:
		Printf("x 是 float64 型")
	case func(int) float64:
		Printf("x 是 func(int) 型")
	case bool, string:
		Printf("x 是 bool 或 string 型")
	default:
		Printf("未知型")
	}
	Println()
	Println()
}
