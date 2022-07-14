/*
指针是实现标准库中flag包的关键技术，它使用命令行参数来设置对应变量的值
，而这些对应命令行标志参数的变量可能会零散分布在整个程序中。为了说明这一点，
在早些的echo版本中，就包含了两个可选的命令行参数：-n用于忽略行尾的换行符
，-s sep用于指定分隔字符（默认是空格）。下面这是第四个版本，对应包路径为gopl.io/ch2/echo4。
*/
package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "忽略行尾的换行符")
var s = flag.String("s", " ", "指定分隔字符（默认是空格）")

func main() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *s))
	if !*n {
		fmt.Println()
	}
}
