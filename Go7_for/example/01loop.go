package example

import (
	. "fmt"
)

func Loop() {
	Println("一、for循环简单使用包含range的使用")
	var b int = 15
	var a int
	numbers := [6]int{1, 3, 5, 7}
	for a := 0; a < 10; a++ {
		Printf("a的值为:%d\n", a)
	}
	for a < b {
		a++
		Printf("a的值为:%d\n", a)
	}

	// for 循环的 range 格式可以对 slice、map、数组、字符串等进行迭代循环。格式如下：
	// for key, value := range oldMap {
	// 	newMap[key] = value
	// }
	for key, val := range numbers {
		Printf("第%d位的值为:%d\n", key, val)
	}
	Println()
}
