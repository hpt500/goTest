package example

import (
	. "fmt"
)

func Factorial() {
	Println("一.递归函数中的阶乘应用")
	var tnum int = 15
	Printf("%d的阶乘是%d\n", tnum, FacFunc(uint64(tnum)))
	Println()
}

func FacFunc(n uint64) uint64 {
	var result uint64
	if n > 0 {
		result = n * FacFunc(n-1)
		return result
	}

	return 1

}
