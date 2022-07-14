package main

import (
	"fmt"
	"time"
)

// 实现channal 的多路复用
func main() {
	intChan := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan <- i
	}

	strChan := make(chan string, 10)
	for i := 0; i < 10; i++ {
		strChan <- fmt.Sprintf("strChan -- %d", i)
	}
	// 开始多路复用 同时使用几个管道
	// 使用select多路复用时不需要关闭管道
	for {
		select {
		case v := <-intChan:
			fmt.Println("Get intChan the Value: ", v)
			time.Sleep(time.Millisecond * 100)
		case v := <-strChan:
			fmt.Println("Get strChan the Value: ", v)
		default:
			fmt.Printf("所有数据 获取完毕\n")
			return
		}
	}
}
