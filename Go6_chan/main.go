package main

import (
	. "fmt"

	"./example"
)

func main() {
	Println("channel的用法")
	// 1.增加缓冲期并休眠1秒钟
	// ch := make(chan int, 1)
	// ch <- 1
	// go func() {
	// 	v := <-ch
	// 	Println(v)
	// }()
	// time.Sleep(1 * time.Second)
	// 2.ch <- 1放置到子线程代码后面
	ch := make(chan int)
	go func() {
		v := <-ch
		Println(v)
	}()
	ch <- 1
	Println("2")
	Println()

	pro := make(chan int)
	go example.produce(pro)

}
