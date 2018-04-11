package example

import (
	. "fmt"
	"time"
)

func Simple1() {
	Println("channel有缓冲的用法")
	// 1.增加缓冲期并休眠1秒钟
	ch := make(chan int, 1)
	ch <- 1
	go func() {
		v := <-ch
		Println(v)
	}()
	time.Sleep(1 * time.Second)
	Println("2")
	Println()
}
func Simple2() {
	Println("channel无缓冲的用法")
	// 2.ch <- 1放置到子线程代码后面
	ch := make(chan int)
	go func() {
		v := <-ch
		Println(v)
	}()
	ch <- 1
	Println("2")
	Println()
}
