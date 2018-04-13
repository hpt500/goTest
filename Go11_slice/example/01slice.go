package example

import (
	. "fmt"
)

func Slice() {
	Println("一.1slice的初步应用")
	// 也可以指定容量，其中capacity为可选参数。
	// make([]type, length, capacity)
	// 这里 len 是数组的长度并且也是切片的初始长度。
	var sli = make([]int, 3, 5)
	PrintSlice(sli)
	Println()

	Println("一.2nil空切片")
	var snil []int
	PrintSlice(snil)
	if snil == nil {
		println("切片snil是空的")
	}
	Println()

	Println("一.3切片截取")
	num := []int{1, 2, 3, 4, 5, 6}
	Println("-1.打印原始切片:")
	PrintSlice(num)
	// 注意:cap指的是指向数组的首元素到数组最后一个元素的个数,否管是否切到最后一个元素,只认准初始元素的位置
	Println("-2.打印切片从索引1(包括)到索引4(不包括):")
	PrintSlice(num[1:4])
	Println("-3.默认下限为0:")
	PrintSlice(num[:3])
	Println("-4.默认上线为len(num):")
	PrintSlice(num[2:])
	Println()

}

func PrintSlice(x []int) {
	Printf("len=%d,cap=%d,slice=%v\n", len(x), cap(x), x)
}
