package operator

import "fmt"

func init() {
	var c int
	c = 200
	// 五、赋值运算符
	println("五、赋值运算符")

	c <<= 2
	fmt.Printf("第一行  - <<= 运算符实例，c 值为 = %d\n", c)

	c >>= 2
	fmt.Printf("第二行 - >>= 运算符实例，c 值为 = %d\n", c)

	// 200的二进制为 11001000   2的二进制为 10   200&2 = 0
	c &= 2 //即为0
	fmt.Printf("第三行 - &= 运算符实例，c 值为 = %d\n", c)

	// 0的二进制为0  2的二进制为 10   0^2 = 2
	c ^= 2
	fmt.Printf("第四行 - ^= 运算符实例，c 值为 = %d\n", c)

	// 2的二进制为10  2的二进制为 10   2|2 = 2
	c |= 2
	fmt.Printf("第五行 - |= 运算符实例，c 值为 = %d\n", c)

	println()

}
