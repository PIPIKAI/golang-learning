package main

import (
	"fmt"
	"gomod_example/calc"

	// 给包起别名
	tl "gomod_example/tools"
	// 导入匿名包
	// _ "gomod_example/tools"
)

func init() {
	fmt.Printf("\"main\": %v\n", a)
}

const a int = 2233

func main() {
	fmt.Printf("calc.Add(10, 2): %v\n", calc.Add(10, 2))
	fmt.Printf("calc.Age: %v\n", calc.Age)
	// 访问私有 报错 .\main.go:12:35: undefined: calc.aa
	// fmt.Printf("calc.aa: %v\n", calc.aa)

	tl.Printaa()
}
