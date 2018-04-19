package example

import (
	. "fmt"
)

func TypeCover() {
	Println("四.go语言类型转换")
	var x, y int = 17, 5
	var mean float32
	mean = float32(x) / float32(y)

	Printf("mean的值为:%f\n", mean)
	Println()

}
