package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

//  写数据
//  定义为 只写
func ch1(ch chan<- int) {
	for i := 1; i < 11; i++ {
		ch <- i
		fmt.Printf("写入数据 %v \n", i)
		time.Sleep(time.Millisecond * 500)
	}
	// 再用range 遍历管道之前需要将管道关闭
	close(ch)
	wg.Done()
}

// 读数据
func ch2(ch <-chan int) {
	for v := range ch {
		fmt.Printf("读数据 %v \n", v)
		time.Sleep(time.Millisecond * 50)
	}
	wg.Done()
}
func main() {
	ch := make(chan int, 10)
	wg.Add(1)
	go ch1(ch)
	wg.Add(1)
	go ch2(ch)
	wg.Wait()
}
