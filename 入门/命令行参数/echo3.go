package main

import (
	"fmt"
	"os"
	// 这里用到了strings 这个包
	"strings"
)

func main() {
    fmt.Println(strings.Join(os.Args[1:], " "))
}