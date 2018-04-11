package example

import (
	"fmt"
	"time"
)

func Nobuffer() {
	// 1.channel无缓冲的情况下
	// ->生产者将值0写入导致子线程阻塞
	// ->直到消费者将值0取出 并输出 receive: 0
	// ->消费者这时再进行循环取出下一个值时因生产者还没有将数据存入导致消费者进入阻塞状态
	// ->回归生产者--阻塞消除-循环恢复-输出 send: 0 -同时进行下一次循环写入1
	// ->这时的消费者刚好因第二次循环取出值导致进入阻塞状态
	// ->因此生产者不产生阻塞并成功写入1 并输出 send: 1 进入下一次循环写入2 进入阻塞
	// ->回归消费者--阻塞消除-循环恢复-取出1并输出 receive: 1-同时进行下一次循环读取2 并输出receive: 2
	// 以此进行循环直到结束
	fmt.Println("实例无缓冲情况")
	pro := make(chan int)
	go produce(pro)
	go consumer(pro)
	time.Sleep(1 * time.Second)
}

func Hasbuffer() {
	// 2.channek有缓冲的情况下
	// ->缓冲区可以存储10个int类型的整
	// ->执行生产者线程的时候,不会阻塞而一次性写入10个数值
	// ->执行消费者线程的时候,同样->一次性读取10个数值
	fmt.Println("实例带缓冲情况")
	pro := make(chan int, 10)
	go produce(pro)
	go consumer(pro)
	time.Sleep(1 * time.Second)
}

func produce(p chan<- int) {
	for i := 0; i < 10; i++ {
		p <- i
		fmt.Println("send:", i)
	}
}
func consumer(c <-chan int) {
	for i := 0; i < 10; i++ {
		v := <-c
		fmt.Println("receive:", v)
	}
}
