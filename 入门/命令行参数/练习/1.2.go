// 练习 1.2： 修改echo程序，使其打印每个参数的索引和值，每个一行。
package main

import (
	"fmt"
	"os"
)
func main() {
	for idx, v := range os.Args[1:] {
		fmt.Printf("index: %d value: %s\n", idx, v)
	}
}