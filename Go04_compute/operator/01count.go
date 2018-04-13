package operator

import (
	"fmt"
)

func init() {
	// 一、算术运算符+-*/%  ++ --
	var a, b, c int = 21, 10, 0
	fmt.Println("一、算术运算符")
	c = a + b
	fmt.Printf("第一行中c的值为:%d\n", c)
	c = a - b
	fmt.Printf("第二行中c的值为:%d\n", c)
	c = a * b
	fmt.Printf("第三行中c的值为:%d\n", c)
	c = a / b
	fmt.Printf("第四行中c的值为:%d\n", c)
	c = a % b
	fmt.Printf("第五行中c的值为:%d\n", c)
	a = 21
	a--
	fmt.Printf("第六行中a的值为:%d\n", a)
	a++
	fmt.Printf("第七行中a的值为:%d\n", a)
	fmt.Println()

}
