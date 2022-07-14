package main

import "fmt"

func main() {
	// 管道
	// 1. 在默认情况下管道是双向的 遵循先入先出的规则
	chans := make(chan int, 10)
	// 表示管道的入队
	for i := 0; i < 10; i++ {
		chans <- i
	}
	// 管道的出队
	<-chans
	// 输出是一个地址 ,他是一个引用类型
	// fmt.Println(chans)
	// 遍历
	for i, length := 0, len(chans); i < length; i++ {
		fmt.Println(<-chans)
	}

	// 2. 声明只写管道
	write_only_chans := make(chan<- int, 2)
	write_only_chans <- 2

	// <-write_only_chans   错误:cannot receive from send-only channel write_only_chans (variable of type chan<- int)
	fmt.Println(write_only_chans)

	// 3. 声明只读管道
	read_only_chans := make(<-chan int, 12)

	fmt.Println(read_only_chans)
}
