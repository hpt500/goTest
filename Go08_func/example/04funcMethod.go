package example

import (
	. "fmt"
)

type Circle struct {
	radius float64
}

func FuncMethod() {
	Println("四、go语言中同时有函数和方法")
	var c1 Circle
	c1.radius = 10.00
	Println("Area of Circle(c1) = ", c1.getArea())
}

func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}
