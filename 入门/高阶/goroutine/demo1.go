package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func test(num int) {
	defer wg.Done()
	for i := 0; i < 3; i++ {
		fmt.Printf("第 %d 个协程打印的第 %d条数据\n", num, i)
		time.Sleep(time.Millisecond * 100)
	}
}
func main() {
	for i := 0; i < 4; i++ {
		wg.Add(1)
		go test(i)
	}
	wg.Wait()
	fmt.Printf("\"主线程结束\": %v\n", "主线程结束")
	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())
}
