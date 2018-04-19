package example

import (
	. "fmt"
)

// 定义一个结构体
type Num struct {
	numOne int
	numTwo int
}

// 实现erros接口
func (n *Num) Error() string {
	strFormat := `
	Cannot proceed, the numTwo is zero.
	numOne: %d
	numTwo: 0
	`
	return Sprintf(strFormat, n.numOne)
}

// 一个int类型除法函数
func divide(num1 int, num2 int) (result int, errorMsg string) {
	if num2 == 0 {
		divi := Num{
			numOne: num1,
			numTwo: num2,
		}
		errorMsg = divi.Error()
		return
	}
	return num1 / num2, ""

}

func InterError() {
	Println("二.go语言的错误处理机制")
	// 正常情况下
	if result, errorMsg := divide(100, 10); errorMsg == "" {
		Println("100/10=", result)
	}
	// 被除数为0的情况下
	if _, errorMsg := divide(100, 0); errorMsg != "" {
		Println(errorMsg)
	}

}
