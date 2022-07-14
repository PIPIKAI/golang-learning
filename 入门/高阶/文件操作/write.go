package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

// 写入文件

// 方法一: read
func fn1() {
	file, err := os.OpenFile("D:\\CODES\\go_codes\\高阶\\文件操作\\writed_file.in", os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0666)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fn1 :%v\n", err)
		return
	}
	file.WriteString("123456789\n")
	file.Write([]byte("987654312\n"))
}

// 方法二: bufio
func fn2() {
	file, err := os.OpenFile("D:\\CODES\\go_codes\\高阶\\文件操作\\writed_file.in", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fn2 :%v\n", err)
		return
	}

	writer := bufio.NewWriter(file)
	writer.WriteString("呵呵呵🙂\n")
	writer.WriteString("哈哈哈🕶\n")

	defer writer.Flush()
}

// 方法三: 使用ioutil

func fn3() {

	err := ioutil.WriteFile("D:\\CODES\\go_codes\\高阶\\文件操作\\writed_file.in", []byte("使用ioutil🤣\n"), 0666)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fn3 :%v\n", err)
		return
	}
}

// 写入文件
func main() {
	fn3()
}
