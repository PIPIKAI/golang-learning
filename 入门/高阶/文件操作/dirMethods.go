package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.MkdirAll("mkdir/mkdir1/mkdir2", 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "no : %v\n", err)
		return
	}

	// 删除目录下所有文件
	_ = os.RemoveAll("mkdir/mkdir1")
	// 删除该目录 ,或者文件
	// _ = os.Remove("mkdir")

	// 重命名
	_ = os.Rename("mkdir", "newmkdir")
}
