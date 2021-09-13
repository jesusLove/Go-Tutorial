package main

import "fmt"

/**
	通道 Channel 是 Go 语言独有的并发编程模式和编程哲学。
	通道用来在多个 goroutine 之间传递数据
	通道类型的值本身是并发安全的。

	通道：相当于先进先出的队列。通道中的元素严格按照发送顺序排列。
	通道中的元素值得发送和接受都是操作符 <- ，称为接送操作符。
*/
func main() {
	// ! 1. 定义通道: 使用 make
	// chan int: chan 表示通道类型的关键字；int 说明该通道元素类型
	// 3 : 通道容量，即通道最多缓存多少元素值。
	ch1 := make(chan int, 3)
	// ! 2. 下通道中发送三个元素；由于容量是3，所以它们被缓存在通道中。
	ch1 <- 2
	ch1 <- 1
	ch1 <- 3

	// ! 3 接受通道中元素
	el1 := <- ch1
	fmt.Printf("The first element: %d\n", el1)
}