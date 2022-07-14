package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// 读取文件

// 方法一: read
func fn1() {
	file, err := os.Open("D:\\CODES\\go_codes\\高阶\\文件操作\\readandwrite.go")
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fn1 :%v\n", err)
		return
	}
	readed := []byte{}
	for {
		temp := make([]byte, 256)
		_, err := file.Read(temp)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "fnf1 :%v\n", err)
			return
		}
		// 注意这一行
		readed = append(readed, temp[:]...)
	}
	fmt.Printf("readed: %v\n", string(readed))
}

// 方法二: bufio
func fn2() {
	file, err := os.Open("D:\\CODES\\go_codes\\高阶\\文件操作\\readandwrite.go")
	defer file.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fn2 :%v\n", err)
		return
	}
	// 用bufio 读取文件
	readed := ""
	reader := bufio.NewReader(file)
	for {
		// 表示一次读取一行
		content, err := reader.ReadString('\n')
		if err == io.EOF {
			// 注意这个地方
			readed += content
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "fn2 :%v\n", err)
			return
		}
		// 注意这一行
		readed += content
	}
	fmt.Printf("readed: \n%v\n", readed)
}

// 方法三: 使用ioutil

func fn3() {
	bytes, err := ioutil.ReadFile("D:\\CODES\\go_codes\\高阶\\文件操作\\readandwrite.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "fn2 :%v\n", err)
		return
	}
	fmt.Println(string(bytes))
}

// 写入文件
func main() {
	fn3()
}
