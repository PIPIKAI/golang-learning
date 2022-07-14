// 练习 1.4： 修改dup2，出现重复的行时打印文件名称。

package main

import (
	"fmt"
	"bufio"
	"os"
)
func countLines(f *os.File, counts map[string]int ) bool {
	have_reapeat_line := false 
	input := bufio.NewScanner(f)
	for input.Scan(){
		counts[input.Text()] ++
		if counts[input.Text()] > 1{
			have_reapeat_line =true
		}
	}
	return have_reapeat_line
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
			pr := countLines(f , counts )
			if pr {
				fmt.Printf("出现了重复的行，其文件名称是 ：%s\n", filename)
			}
		}
	}
	// NOTE: ignoring potential errors from input.Err()
	for line,n := range counts {
		if n > 1{
			fmt.Printf("%d\t%s\n", n,line)
		}
	}
}