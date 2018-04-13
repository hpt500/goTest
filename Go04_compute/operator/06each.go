package operator

import "fmt"

func init() {
	// 六、其他运算符
	var a int = 2
	var b int32
	var ptr *int
	println("其他运算符")

	// 运算符实例
	fmt.Printf("第一行a变量的类型为:%T\n", a)
	fmt.Printf("第二行b变量的类型为:%T\n", b)

	// 1.&返回变量存储地址 即&a将给出变量的实际地址
	// 2.*指针变量 即*a是一个指针变量
	println("变量a的存储地址为:", &a)
	ptr = &a // ptr包含变量a的地址
	println("*ptr的值为:", *ptr)
	println()
}
