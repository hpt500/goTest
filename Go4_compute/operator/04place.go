package operator

func init() {
	var a, b int = 60, 13
	// 四、位运算符
	println("四、位运算符")
	// & 按位与运算符"&"是双目运算符。 即取两数二进制数值对应部分。
	println("第一行输出a&b的值:", a&b)
	// | 按位或运算符"|"是双目运算符。 即混合两数二进制数值所有部分
	println("第二行输出a|b的值:", a|b)
	// ^ 按位异或运算符"^"是双目运算符。 即混合除两数二进制数值相同部分之外的部分。
	println("第三行输出a^b的值:", a^b)
	// << 左移运算符"<<"是双目运算符。 即在第一个数值的二进制数值基础上左移第二个数值的位数
	println("第四行输出a<<2的值:", a<<2)
	// >> 右移运算符">>"是双目运算符。 同上相反
	println("第五航输出a>>2的值:", a>>2)
	println()

}
