package example

import (
	. "fmt"
)

func AppendCopy() {
	Println("二.append()函数的使用")
	var num []int
	Println("-1.打印原始切片")
	PrintSlice(num)
	Println("-2.允许切片追加空切片")
	num = append(num, 0)
	PrintSlice(num)
	Println("-3.向切片追加一个元素")
	num = append(num, 1)
	PrintSlice(num)
	Println("-4.向切片添加多个元素")
	num = append(num, 2, 3, 4, 5)
	PrintSlice(num)
	Println("-4.创建切片是num的两倍容量")
	num2 := make([]int, len(num), (cap(num) * 2))
	copy(num2, num)
	PrintSlice(num2)
	Println()

}
