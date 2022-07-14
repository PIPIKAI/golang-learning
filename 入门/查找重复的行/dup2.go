// 对文件做拷贝、打印、搜索、排序、统计或类似事情的程序都有一个差不多的程序结构：一个处理输入的循环，
// 在每个元素上执行计算处理，在处理的同时或最后产生输出。我们会展示一个名为dup的程序的三个版本；
// 灵感来自于Unix的uniq命令，其寻找相邻的重复行。该程序使用的结构和包是个参考范例，可以方便地修改。

package main

import (
	"fmt"
	"bufio"
	"os"
)
func countLines(f *os.File, counts map[string]int)  {
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()] ++
	}

}

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) ==0{
		countLines(os.Stdin,counts)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr ,"dup2: %v\n",err)
				continue
			}
			countLines(f , counts)
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	for line,n := range counts {
		if n > 1{
			fmt.Printf("%d\t%s\n", n,line)
		}
	}
}