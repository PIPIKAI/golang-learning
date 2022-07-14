package main

import (
	"fmt"
	"sync"
	"time"
)

// 统计素数

const num int = 4
const N int = 120000

var k int = N / num
var res = []int{}
var wg sync.WaitGroup

func solve(th int) {
	defer wg.Done()
	for i := th * k; i < (th+1)*k; i++ {
		if i < 2 {
			continue
		}
		var flag bool = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				flag = false
				break
			}
		}
		if flag {
			res = append(res, i)
		}
	}
}
func main() {

	start := time.Now()

	for i := 0; i < num; i++ {
		wg.Add(1)
		solve(i)
	}
	wg.Wait()
	fmt.Printf("运行时间 %v s\n", time.Now().Sub(start).Seconds())
	fmt.Println(res)
}
