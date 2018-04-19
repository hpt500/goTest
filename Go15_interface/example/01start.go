package example

import (
	. "fmt"
)

type phone interface {
	call()
}

type TomCall struct {
}

type TinaCall struct {
}

func (name TomCall) call() {
	Println("I am Tom,i can call you")
}

func (name TinaCall) call() {
	Println("I am Tina,i can call you")
}

func Start() {
	Println("一.go语言接口的创建")
	var ph phone
	// new(struct_name) 进行该结构体的赋值
	ph = new(TomCall)
	ph.call()

	ph = new(TinaCall)
	ph.call()

	Println()

}
