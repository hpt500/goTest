package main

import (
	. "fmt"

	"./example"
)

func main() {
	Println()
	// 一、for循环简单使用包含range的使用
	example.Loop()
	// 二、循环嵌套->输出100以内的素数
	example.Nest()
	// 三、break语句的使用->跳出当前循环
	example.Break()
	// 四、continue语句的使用->跳出当前循环执行下一次循环
	example.Continue()
	// 五、goto语句的使用->可用来实现条件转移,构成循环,跳出循环体等功能
	example.Goto()
	// 六、goto语句的使用2
	example.GotoExa()
}
