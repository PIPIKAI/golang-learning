package main

import "fmt"

// 结构体，当首字母命名为大写的时候表示公有，首字母为小写的时候表示私用
// 结构体的嵌套可以实现类的继承

type Dog struct {
	age       int
	hairColor string
	// 添加一个匿名的结构体 当调用匿名结构体中的变量时，若该名称是独一无二的则可以省略结构体字段
	Animals
}
type Animals struct {
	Name string
}

// 结构体的函数的实现 func (s StructName) 函数名称 (形参) (返回值) {}
func (self Animals) run() {
	fmt.Printf("self.name: %v\n", self.Name)
}
func (self Animals) init() {
	fmt.Printf("Animals init")
}
func (self *Animals) rename(s string) {
	self.Name = s
}

func main() {
	dog1 := Dog{
		age:       1,
		hairColor: "red",
		Animals: Animals{
			Name: "cao",
		},
	}

	dog1.rename("dog1")

	fmt.Printf("dog1: %#v\n", dog1)
}
