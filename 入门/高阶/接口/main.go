package main

import (
	"fmt"
)

// 接口的定义
type Usber interface {
}

func main() {
	// 空接口的使用
	// 表示可以传入任意类型
	mp := map[string]interface{}{
		"name": "ally",
		"age":  20,
	}
	fmt.Printf("mp: %v  %T\n", mp, mp)
	// 类型断言 .(type)
	// 返回两个值，第二个是bool用来判断 ,若 true 则返回正确的类型
	if v, ok := mp["name"].(int); ok {
		fmt.Printf("v : %v type %T  ", v, v)
	} else {
		fmt.Printf("\"不是string这个类型\": %v\n", "不是 int 这个类型")
		fmt.Printf("v : %v type %T  \n", v, v)
	}
	// 也可以这样用来判断数据的类型
	switch v := mp["age"].(type) {
	case int:
		fmt.Printf("是 int 这个类型 value is %v", v)
	case string:
		fmt.Printf("是 string 这个类型 value is %v", v)
	default:
		fmt.Printf("err")

	}
	// 可以用接口来嵌套接口
	// 也可以一个结构体 接受多个接口
}
